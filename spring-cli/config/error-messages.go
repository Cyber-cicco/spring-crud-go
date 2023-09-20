package config

/*
*   Liste des différents messages d'erreur possibles
*/

var ERR_NOT_ENOUGH_ARGS_MAIN = "Error: the number of arguments must be greater than 2"
var ERR_JPA_ARGS = "Error: you must specify the entity name and its fields to use the jpa command"
var ERR_CLASS_ARGS = "Error: class name is mandatory"
var ERR_BAD_ARGS = "Incorrect usage of the command"
var ERR_TYPE_NOT_GIVEN = "The character ':' is used to specify the type. What follows this character must not be empty"
var ERR_FILE_CREATION = "Error in creating a file"
var ERR_DIR_CREATION = "Error in creating a directory"
var ERR_MARSHARL = "Error in deserializing an object from JSON"
var ERR_OPEN_CONFIG = "Error opening the configuration file. If it doesn't exist, please use the init command to create it.\nOtherwise, make sure to run the command in a directory containing a valid configuration file."
var ERR_UNMARSHAL = "Error in creating an object from JSON"
var ERR_BAD_TS_DIRECTORY = "Error: you need to specify a directory to write the TypeScript files in the spring-parameters.json configuration file"
var ERR_BAD_CONFIG_PACKAGE = "Error: the package specified in the configuration file does not seem to point to an existing directory."
var ERR_JPA_DIR_OPEN = "Error: the folder supposed to contain the JPA configuration files does not exist. Remember to create them using the 'jpa' command before running the 'pr' command."
var ERR_CURRENT_DIR_OPEN = "Error: opening the root directory of the project seems impossible"
var ERR_JPA_FILE_READ = "Error reading the configuration file of a JPA entity"
var ERR_NO_JPA_FILE = "There is no JPA entity configuration file in the jpa folder. Please generate the configuration files using the jpa command before using the spring command"
var ERR_CONFIG_BAD_USAGE = "Error in using the 'config' command"
var ERR_SRC_FOLDER_NOT_IN_ROOT = "The src folder in which all the "
var ERR_JAVA_PARSING_FAILED = "Error parsing a java file"
var ERR_JAVA_ANALYSING = "Error analyzing a java file: a controller method seems not to have been annotated in a way that allows detection"
var ERR_NO_JAVA = "It seems like you're trying to initialize spring-go in a folder that doesn't contain a Java project.\nMake sure the './src/main/java/' folder exists in your directory and try again"
var ERR_TEMPLATE_FILE_READ = "Error reading a template file"

