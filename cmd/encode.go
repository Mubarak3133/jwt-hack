package cmd

import (
	"encoding/json"
	"fmt"

	jwtInterface "github.com/hahwul/jwt-hack/pkg/jwt"
	"github.com/spf13/cobra"
)

var secret, algo string

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode [JSON]",
	Short: "Encode json to JWT",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			mapInterface := []byte(args[0])
			var raw map[string]interface{}
			if err := json.Unmarshal(mapInterface, &raw); err != nil {
				// err
			}
			fmt.Println(jwtInterface.JWTencode(raw, secret, algo))

		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.PersistentFlags().StringVar(&secret, "secret", "", "secret key for JWT signature")
	encodeCmd.PersistentFlags().StringVar(&algo, "algorithm", "HS256", "Algorithm of JWT\ne.g) HS/RS/ECDA - 256,384,512")

}
