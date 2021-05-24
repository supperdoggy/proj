package sctructs

const ApiV1 = "/api/v1"
const TokenExpirationHours = 64

// took from here https://stackoverflow.com/questions/201323/how-to-validate-an-email-address-using-a-regular-expression
const RegexpEmail = "([-!#-'*+/-9=?A-Z^-~]+(\\.[-!#-'*+/-9=?A-Z^-~]+)*|\"([]!#-[^-~ \\t]|(\\\\[\\t -~]))+\")@[0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?(\\.[0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?)+"
