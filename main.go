package main;

import "myApp/logger";

func main() {
	logger.Logger(logger.Types{ERROR: "ERROR"},"The Error Has Occurd In Line 64");
	logger.Logger(logger.Types{WARN: "WARN"},"The Warn Has Occurd In Line 102");
	logger.Logger(logger.Types{LOG: "LOG"},"The LOG Has Occurd In Line 12");
}