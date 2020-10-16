import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:client/models/topic.dart';
import 'package:client/topic/repository/topic_repository.dart';
import 'package:client/topic/services/get_topic_service.dart';

class HomeController {
  FirebaseFirestore firestore;

  HomeController() {
    this.firestore = FirebaseFirestore.instance;
  }

  Stream<List<Topic>> getTopics() {
    TopicRepository topicRepository =
        TopicRepository(firestore: this.firestore);
    GetTopicService getTopicService =
        GetTopicService(tpcRepository: topicRepository);

    return getTopicService.execute();
  }
}
