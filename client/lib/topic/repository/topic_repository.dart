import 'package:cloud_firestore/cloud_firestore.dart';

class TopicRepository {
  final FirebaseFirestore firestore;

  TopicRepository({this.firestore});

  Stream<QuerySnapshot> getTopics() {
    return firestore.collection('topics').orderBy('insertedAt', descending: true).snapshots();
  }
}
