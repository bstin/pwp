pwp
===

Password Paste: A simple command-line password management system (written in go)

What does it do for me?
=======================
		
  Simply put, it allows you to quickly enter and retrieve passwords and other secure notes. When it retrieves a password it automatically puts it in the clipboard buffer for 30 secs so you can simply paste it where you need. 

  After 30 secs your clipboard buffer will be restored to whatever information it originally contained. 

  Your password file always stored encryped at: ~/.pwp/pwp.enc

  This file is only decypted in memory and decrytped portions are never written to a disk (unless some OS-level swap event happens), therefore should be reasonably secure even on multi-user systems.

  This is a simple utility, not intended to be a complete solution.  

  The configuration file is stored in ~/.pwp/pwp.conf

Features:
==========
  * Works on linux / mac / solaris
  * Simple operation
  * Copies passwords / notes to clipboard so they can be pasted
  * Auto-resets clipboard back to default state after 30 secs to prevent any snooping
  * Can sync encrypted password files to/from remote sites using scp, so you can use one db file on multiple macs
  

Requirements: 
=============
  * some form of command-line clipboard management tools (xclip linux / pbpaste pbcopy on mac) 


Example usage:
==============
		

		Create a new entry
		===================
		$pwp set google
		Enter username: joe@gmail.com
		Enter password or leave blank for random: ******

	
		Get existing entry
		===================
		$pwp get google
		Enter username (or leave blank): joe@gmail.com
		Password is in buffer for 30 seconds.....
		(30 secs later)
		Buffer reset


		List all entries (no passwords shown)
		===================
		$pwp list

		Edit the entries directly
		=========================
		$pwp edit

		Configure defaults
		===================
		$pwp config

		Push / Pull pwp database to remote site
		================================
		$pwp push 	#(will push to remote site as defined in config)
		$pwp pull 	#(will copy from remote site to local)
	
		
		(note: this requires you to have a ssh id_rsa key setup on both machines - which is outside the scope of this document to explain)

Installation Instructions
=========================
  Coming Soon		

Credits
========

  Original idea by tr0lltherapy of reddit. See thread here: 

  http://www.reddit.com/r/programming/comments/x0nzz/very_simple_nononsense_password_manager_that_uses/c5i7ekn

  I also implemented this as a bash script which can be found under github.com/jposse/pwp

