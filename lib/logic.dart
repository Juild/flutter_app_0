import 'package:http/http.dart' as http;

void authRequest() async {
  var url = Uri.http('localhost:80', '/');
  var response = await http.get(url);
  print(response.body);
}
