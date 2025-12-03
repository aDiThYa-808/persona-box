/*
Package jwtx handles the creation and validation of Personabox's access token(JWT).
It exposes the following four functions for generating and validating tokens:

generator.go: GenerateAccessToken() takes a signing key and claims to create a signed JWT.

validator.go: ValidateAccessToken() takes an access token and verifies if it is valid and issued by Personabox.

validator.go: ParseGoogleJWT() parses the provided Google JWT and populates the provided claims struct with valid claims.

verifyclaims.go: VerifyGoogleClaims() verifies the provided Google JWT claims and returns an error if it is invalid.
*/
package jwtx
