/*
Package jwtx handles the creation and validation of Personabox's access token(JWT).
It exposes the following two functions for generating and validating tokens:
GenerateAccessToken() takes a signing key and claims to create a signed JWT;
ValidateAccessToken() takes an access token and verifies if it is valid and issued by Personabox.
Personabox's access token is signed using HS256.
*/
package jwtx
