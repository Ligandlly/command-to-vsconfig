grammar Command;

command: bin file arg+;

bin: String;
file: String;
arg: Key value | Key '=' value | Key;

Key: ('-' | '--') String;
value: String;

Whitespace: [ \t\n]+ -> skip;
String: ([0-9a-zA-Z_.:/"'-] | '[' | ']')+;

// PYTHON: 'python'; PYTHON3: 'python3';