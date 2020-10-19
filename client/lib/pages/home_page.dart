import 'package:flutter/material.dart';
import 'package:client/pages/home_controller.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage>
    with SingleTickerProviderStateMixin {
  AnimationController _controller;
  HomeController hc;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(vsync: this);
    hc = HomeController();
  }

  @override
  void dispose() {
    super.dispose();
    _controller.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return StreamBuilder(
      stream: this.hc.getTopics(),
      builder: (context, snapshot) {
        print(snapshot);
        if (snapshot.hasData) {
          return ListView.builder(
            itemCount: snapshot.data.length,
            itemBuilder: (context, index) {
              var topic = snapshot.data[index];
              return ListTile(
                  title: Text("${topic.description}  ${topic.payload}"),
                  leading: Icon(Icons.colorize));
            },
          );
        } else {
          return _loadingIndicator();
        }
      },
    );
  }
}

_loadingIndicator() {
  return Center(
    child: CircularProgressIndicator(),
  );
}
