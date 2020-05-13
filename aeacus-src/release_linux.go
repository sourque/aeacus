package main

import (
	"fmt"
	"strconv"
)

func writeDesktopFiles(mc *metaConfig) {
	if mc.Cli.Bool("v") {
		infoPrint("Writing shortcuts to Desktop...")
	}
	shellCommand("cp " + mc.DirPath + "misc/*.desktop /home/" + mc.Config.User + "/Desktop/")
	shellCommand("chmod +x /home/" + mc.Config.User + "/Desktop/*.desktop")
	shellCommand("chown " + mc.Config.User + ":" + mc.Config.User + " /home/" + mc.Config.User + "/Desktop/*")

	if mc.Cli.Bool("v") {
		infoPrint("Writing FQ files to Desktop...")
	}

	var numFQ int
	prompt("How many FQs would you like? ")
	fmt.Scanln(&numFQ)

	var fqAns []string

	counter := 1
	for counter <= numFQ {
		fileName := "FQ" + strconv.Itoa(counter)
		var fqQ string
		prompt(fileName + " - " + "Q: ")
		fmt.Scanln(&fqQ)

		var fqA string
		prompt(fileName + " - " + "A: ")
		fmt.Scanln(&fqA)

		fqAns = append(fqAns, fileName+" - "+fqA)
		shellCommand("echo -e 'Q: " + fqQ + "\nA: ' > " + "/home/" + mc.Config.User + "/Desktop/" + fileName + ".txt")
		if mc.Cli.Bool("v") {
			infoPrint("Wrote " + fileName + ".txt to Desktop")
		}
		counter++
	}
	infoPrint("Remember to add the following to your scoring.conf!")

	for _, i := range fqAns {
		fmt.Printf("%s\n", i)
	}
}

func installService(mc *metaConfig) {
	if mc.Cli.Bool("v") {
		infoPrint("Installing service...")
	}
	shellCommand("cp /opt/aeacus/misc/aeacus-client /etc/init.d/")
	shellCommand("chmod +x /etc/init.d/aeacus-client")
	shellCommand("systemctl enable aeacus-client")
}

func cleanUp(mc *metaConfig) {

	if mc.Cli.Bool("v") {
		infoPrint("Changing perms to 755 in /opt/aeacus...")
	}
	shellCommand("chmod 755 -R /opt/aeacus")

	if mc.Cli.Bool("v") {
		infoPrint("Removing .viminfo files...")
	}
	shellCommand("find / -name \".viminfo\" -delete")

	if mc.Cli.Bool("v") {
		infoPrint("Symlinking .bash_history and .zsh_history to /dev/null...")
	}
	shellCommand("find / -name \".bash_history\" -exec ln -sf /dev/null {} \\;")
	shellCommand("find / -name \".zsh_history\" -exec ln -sf /dev/null {} \\;")

	if mc.Cli.Bool("v") {
		infoPrint("Removing .swp files")
	}
	shellCommand("find / -type f -iname '*.swp' -delete")

	if mc.Cli.Bool("v") {
		infoPrint("Removing .local files")
	}
	shellCommand("rm -rf /root/.local /home/*/.local/")

	if mc.Cli.Bool("v") {
		infoPrint("Removing cache...")
	}
	shellCommand("rm -rf /root/.cache /home/*/.cache/")

	if mc.Cli.Bool("v") {
		infoPrint("Removing swap and temp Desktop files")
	}
	shellCommand("rm -rf /home/*/Desktop/*~")

	if mc.Cli.Bool("v") {
		infoPrint("Removing crash and VMWare data...")
	}
	shellCommand("rm -f /var/VMwareDnD/* /var/crash/*.crash")

	if mc.Cli.Bool("v") {
		infoPrint("Removing apt and dpkg logs...")
	}
	shellCommand("rm -rf /var/log/apt/* /var/log/dpkg.log")

	if mc.Cli.Bool("v") {
		infoPrint("Removing auth and syslog")
	}
	shellCommand("rm -f /var/log/auth.log* /var/log/syslog*")

	if mc.Cli.Bool("v") {
		infoPrint("Removing initial package list")
	}
	shellCommand("rm -f /var/log/installer/initial-status.gz")

	if mc.Cli.Bool("v") {
		infoPrint("Removing scoring.conf...")
	}
	shellCommand("rm /opt/aeacus/scoring.conf*")

	if mc.Cli.Bool("v") {
		infoPrint("Removing other setup files...")
	}
	shellCommand("rm -rf /opt/aeacus/misc /opt/aeacus/web/assets/previous.txt /opt/aeacus/ReadMe.conf /opt/aeacus/README.md /opt/aeacus/TODO.md")

	if mc.Cli.Bool("v") {
		infoPrint("Removing aeacus binary...")
	}
	shellCommand("rm /opt/aeacus/aeacus")
}
