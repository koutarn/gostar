/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func execInit(projectName string){

	// go.modファイルがあれば処理終了
	if _ , err := os.Stat("go.mod"); err == nil {
		fmt.Println("exits go.mod")
		return
	}

	// go mod initを実行
	// 	ユーザー名、プロジェクト名をいれたら実行してくれる
	fmt.Println("go mod initを実行します")
	fmt.Print("usrename:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userName := scanner.Text()

	out,err := exec.Command("go mod init Github.com." + userName + "/" + projectName,"").Output()
	fmt.Println(out)
	if err != nil {
		fmt.Println("Error:",err)
	}
}

func Create(){

	dir ,err := os.Getwd()
	if err != nil {
		fmt.Println("Error:",err)
	}

	// 	プロジェクト名はフォルダ名から自動取得
	// 	TODO:指定フラグも作る
	projectName := filepath.Base(dir)

	execInit(projectName)

	

	// cobla-cliをインストール
	// exec.Command("").Run()
	// if err != nil {
	// 	fmt.Println("Error:",err)
	// }

	// cobla-initを実行
	exec.Command("coble-cli init").Run()
	if err != nil {
		fmt.Println("Error:",err)
	}

	// cobla add versionを実行
	exec.Command("cobla-cli add version").Run()
	if err != nil {
		fmt.Println("Error:",err)
	}

    // 	Makefileをコピー
    // 	testフォルダを作成



}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gostar",
	Short: "いい感じにCLI開発環境を作成します",
	Long: `GoのCLI開発環境を作成します。`,
	Run: func(cmd *cobra.Command, args []string) { 
		// 作成
		Create()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


