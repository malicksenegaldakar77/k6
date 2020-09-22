/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2020 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package netext

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/lib/testutils/mockresolver"
)

func TestResolver(t *testing.T) {
	t.Parallel()

	host := "myhost"
	mr := mockresolver.New(map[string][]net.IP{
		host: {
			net.ParseIP("127.0.0.10"),
			net.ParseIP("127.0.0.11"),
			net.ParseIP("127.0.0.12"),
		},
	}, nil)

	t.Run("LookupIP", func(t *testing.T) {
		testCases := []struct {
			ttl      time.Duration
			strategy lib.DNSStrategy
			expIP    []net.IP
		}{
			{0, lib.DNSFirst, []net.IP{net.ParseIP("127.0.0.10")}},
			{time.Second, lib.DNSFirst, []net.IP{net.ParseIP("127.0.0.10")}},
			{0, lib.DNSRoundRobin, []net.IP{
				net.ParseIP("127.0.0.10"),
				net.ParseIP("127.0.0.11"),
				net.ParseIP("127.0.0.12"),
				net.ParseIP("127.0.0.10"),
			}},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(fmt.Sprintf("%s_%s", tc.ttl, tc.strategy), func(t *testing.T) {
				r := NewResolver(mr.LookupIPAll, tc.ttl, tc.strategy)
				ip, err := r.LookupIP(host)
				require.NoError(t, err)
				assert.Equal(t, tc.expIP[0], ip)

				if tc.ttl > 0 {
					require.IsType(t, &cacheResolver{}, r)
					cr := r.(*cacheResolver)
					assert.Len(t, cr.cache, 1)
					assert.Equal(t, tc.ttl, cr.ttl)
					firstLookup := cr.cache[host].validTo
					time.Sleep(cr.ttl + 100*time.Millisecond)
					_, err = r.LookupIP(host)
					require.NoError(t, err)
					assert.True(t, cr.cache[host].validTo.After(firstLookup))
				}

				if tc.strategy == lib.DNSRoundRobin {
					ips := []net.IP{ip}
					for i := 0; i < 3; i++ {
						ip, err = r.LookupIP(host)
						require.NoError(t, err)
						ips = append(ips, ip)
					}
					assert.Equal(t, tc.expIP, ips)
				}
			})
		}
	})
}