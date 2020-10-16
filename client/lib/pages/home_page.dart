import 'package:flutter/material.dart';
import 'package:client/pages/home_controller.dart';
import 'package:client/models/topic.dart';

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
        if (snapshot.hasData) {
          return ListView.builder(
            itemCount: snapshot.data.length,
            itemBuilder: (context, index) {
              var topic = snapshot.data[index];

              return ListTile(
                title: Text("${topic.payload}"),
              );
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
