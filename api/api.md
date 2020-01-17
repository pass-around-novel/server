# Pass Around Novel API

## Authentication

```
create_user(username, password, email string, phone string?) token string
login(username, password string) token string
logout(token string) void
forgot_pass(id string) void
reset_pass(token, password string) token string
verify_email(token string) void
verify_phone(token string) void
authenticate(token string) void
change_email(email string) void
change_phone(phone string) void
change_pass(oldpass, newpass string) void
```

## Friends View

```
get_friends(void) friends []{name, username, pic string, since long}
get_requests(outgoing bool) requests []{username, pic string, since long}
add_friend(id string) void
remove_friend(username string) void
accept_friend(username string) void
deny_friend(username string) void
rename_friend(username, name string) void
change_profile_picture(pic []byte) void
```

## Bookshelf View

```
get_books(filter int) books []{id, name, pic string, num_chapters int, since, last_turn long, public, notify bool, turn string, authors []{name, pic string}}
star_book(id string) void
unstar_book(id string) void
finish_book(id string) void
change_book_picture(id string, pic []byte) void
create_book(name string, turn_min, turn_max, turn_type int) id string
rename_book(id, name string) void
share_book(id, user string) void
open_book(id string) void
```

## Novel View

```
novel_info(void) name string, chapters []string, turn_min, turn_max, turn_type int
get_pos(void) chapter, pos int
set_pos(chapter, pos int) void
read_chapter(chapter int) paragraphs []string
get_authors(chapter int) authors []{name, pic string}
get_turn(void) name string?
write(text string) void
end_part(void) void
end_paragraph(void) void
end_chapter(void) void
rename_chapter(chapter int, name string) void
close_book(void) void
```

## Miscellaneous

```
get_image(id string) []byte
notify(void) notifications []{type string, ...}
```

## Notifications

```
email_verify: {}
phone_verify: {}
new_friend_req: {username string}
friend_accept: {username string}
friend_deny: {username string}
friend_removed: {username string}
book_added: {name string}
book_finished: {name string}
book_shared: {name string}
turn: {name string?}
typing: {text []string}
```
