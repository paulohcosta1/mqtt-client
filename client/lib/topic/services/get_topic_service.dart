import 'package:client/topic/repository/topic_repository.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:client/models/topic.dart';

class GetTopicService {
  final TopicRepository tpcRepository;

  GetTopicService({this.tpcRepository});

  Stream<List<Topic>> execute() {
    Stream<QuerySnapshot> tpcSnapshots = tpcRepository.getTopics();
    
    return tpcSnapshots.map((qShot) => qShot.docs.map(
        (item) => Topic(payload: item['payload'])
      ).toList()
    );
  }
}
