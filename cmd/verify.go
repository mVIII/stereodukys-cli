/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Aliases: []string{"vr"},
	Short: "LEGIT CHECK",
	Long: `

	./steroiverify --signature <signature> --message <message>
	
	~OPOU <signature> antikatestise thn ypoografh pou miazei kapws etsi:

	b3bdab1031a7dbb99f77a4e877b3c514d30225d04c88ba876d26b27879ad099ba9440aab278b5d3a30f3fbbd68d68d79ee314e7f41fbd4c870e9d2cb588a0a01
	
	~OPOU <message> antikatestise to minima px:
	
	Kostas edw

	EXAMPLE:
	
	./steroiverify verify --signature b3bdab1031a7dbb99f77a4e877b3c514d30225d04c88ba876d26b27879ad099ba9440aab278b5d3a30f3fbbd68d68d79ee314e7f41fbd4c870e9d2cb588a0a01 --message "Kostas edw"
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		signature,_:=cmd.Flags().GetString("signature")
		message,_:=cmd.Flags().GetString("message")

		if signature==""{
			fmt.Println("Signature required!")
		}else if message==""{
			fmt.Println("Message required!")
		}else{
			verifyMessage(message,signature)
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().String("signature","","Signature to compare!s")
	verifyCmd.Flags().String("message","","Message to verify!")
}

func verifyMessage(message string, signature string){
	edpubsteroidukey := "2db9b63d1576779bae1d20f7ad6f1c0e7c92aa55ad2a0f5b7c1bc0ed8c8bfad8"

	mesb := []byte(message)
	sigb, _ := hex.DecodeString(signature)
	pubb, _ := hex.DecodeString(edpubsteroidukey)

	result:=ed25519.Verify(pubb, mesb, sigb)

	switch result {
	case true:
		fmt.Println("LEGIT")
	case false:
		fmt.Println("FAKE >:(")
	}

}
