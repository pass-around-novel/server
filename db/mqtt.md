# Pass Around Novel MQTT Schema

 Notification Name  | Topic Name                | Payload
--------------------|---------------------------|-------------------
 email_verify       | /{user_uuid}/ve           | void
 phone_verify       | /{user_uuid}/vp           | void
 new_friend_req     | /{user_uuid}/fr           | requester id
 friend_accept      | /{user_uuid}/fa           | receiver id
 friend_deny        | /{user_uuid}/fd           | receiver id
 friend_removed     | /{user_uuid}/fx           | friend id
 book_added         | /{user_uuid}/ba           | book id
 book_finished      | /{user_uuid}/bf           | book id
 book_shared        | /{user_uuid}/bs           | book id
 turn               | /{book_uuid}/{user_uuid}  | void
 typing             | /{book_uuid}              | text
