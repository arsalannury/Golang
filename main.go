package main

import "myApp/logger"

func main() {
	logger.Logger(logger.Types{ERROR: "ERROR"},"The Error Has Occurd In Line 64");
}