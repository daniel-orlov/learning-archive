import 'dart:io';

void main() {
  stdout.writeln('Who would you like to say hello to?');
  String? name = stdin.readLineSync();
  return hello(name!);
}

void hello(String name) {
  print("Hello, $name!");
}