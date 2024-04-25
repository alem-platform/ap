/* ************************************************************************** */
/*                                                                            */
/*                                                        :::     :::::::::   */
/*   ap.go                                             :+: :+:   :+:    :+:   */
/*                                                   +:+   +:+  +:+    +:+    */
/*   By: atlekbai <tlekbai@doodocs.kz>             +#++:++#++: +#++:++#+      */
/*                                                +#+     +#+ +#+             */
/*   Created: 2024/04/24 13:00:17 by atlekbai    #+#     #+# #+#              */
/*   Updated: 2024/04/24 13:00:16 by atlekbai   ###     ### ###               */
/*                                                                            */
/* ************************************************************************** */

package ap_test

import (
	"errors"
	"testing"

	"github.com/alem-platform/ap"
)

func TestPutRune(t *testing.T) {
	type args struct {
		b rune
	}
	tests := []struct {
		name   string
		args   args
		exp    int
		expErr error
	}{
		{
			name: "UTF8",
			args: args{
				b: '💡',
			},
			exp:    4,
			expErr: nil,
		},
		{
			name: "ASCII",
			args: args{
				b: 'A',
			},
			exp:    1,
			expErr: nil,
		},
		{
			name: "ASCII",
			args: args{
				b: '9',
			},
			exp:    1,
			expErr: nil,
		},
		{
			name: "Invalid",
			args: args{
				b: -1,
			},
			exp:    0,
			expErr: ap.ErrInvalidRune,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ap.PutRune(tt.args.b)
			if !errors.Is(err, tt.expErr) {
				t.Errorf("PutRune() error = %v, expErr %v", err, tt.expErr)
				return
			}
			if got != tt.exp {
				t.Errorf("PutRune() = %v, exp %v", got, tt.exp)
			}
		})
	}
}

func ExamplePutRune() {
	ap.PutRune('H')
	ap.PutRune('i')
	// Output: Hi
}
