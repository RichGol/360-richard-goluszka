INFORMATIONAL
=============
Author: Rich Goluszka
Last Updated: 2/16/2021
Project: Coding Standards Validator - Part 1
Course: Applied Programming Languages (CPSC360-1) with Professor Eric Pogue

Contact
-------
Please direct any and all comments/concerns/inquiries to richardjgoluszka@lewisu.edu

ORIGINALITY
===========
Credit to Chapter 8 of _Introducing Go_ by Caleb Doxsey for the code structure to open files 
	and directories using the io/ioutil package.

Credit to [Open Source Initiative](opensource.org/licenses/MIT) for the standard contents of an 
	MIT License.

All other code is the original work of the author and may be used in accordance with the 
	specifications laid out in the LICENSE file.

BUILD / EXECUTE / DEPENDENCY
============================
Required files
--------------
coding-standards-validator-part-1
	-val.go
	-directorychk directory  
		-dirCheck.go
	-licensechk directory  
		-licenseCheck.go  
	-linefmtchk directory  
		-lineFmtCheck.go  
	-utf8chk directory  
		-utf8Check.go  

_Note: The GitHub repository https://github.com/RichGol/360-richard-goluszka contains all_  
_required files plus README.md and LICENSE files. This repository is private and you will *NOT*_  
_be able to access it if you are not an invited collaborator._

Build instructions
------------------
To compile an executable:
1. Open the command-line or terminal
2. Navigate to .../go/src/360-richard-goluszka/sprint-2/coding-standards-validator-part-1
3. Run `go install` within each subdirectory (directorychk / licensechk / linefmtchk / utf8chk)
4. Run `go build` within the val subdirectory (.../coding-standards-validator-part-1/val)
You should now have a `val.exe` executable to call in order to run the program

Execution instructions
----------------------
1. Build the program _(using above instructions)_
2. Run val.exe and specify path to project when prompted
3. Optionally use `val.exe detail` to view detailed validation information
4. Optionally use `val.exe help` to view help instructions
