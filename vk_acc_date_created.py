# coding: utf8


import requests
import vk_api


def authorization():
    access_token = "TOKEN" # вставить свой

    vk_session = vk_api.VkApi(token=access_token)
    vk_session._auth_token()

    return vk_session


global VK_SESSION
VK_SESSION = authorization()


def get_profile_url():
    profile_url = raw_input("Profile URL: ")
    return profile_url


def select_screenname(profile_url):
    screenname = profile_url.replace("https://vk.com/", "")
    return screenname


def get_user_id(screenname):
    values = {
        "user_ids": screenname
    }
    response = VK_SESSION.method("users.get", values)
    user_id = response[0]["id"]
    return user_id


def load_html(user_id):
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0"
    }
    request_url = "https://vk.com/foaf.php?id=" + str(user_id)
    html_response = requests.get(request_url, headers=headers)
    html_text = ''.join(html_response)
    return html_text


def get_name(text):
    begin = text.find("<foaf:name>")
    begin += 11
    end = text.find("</foaf:name>", begin)
    name = text[begin:end]
    return name


def get_acc_date_created(text):
    begin = text.find("<ya:created dc:date=")
    begin += 21
    end = text.find(">", begin)
    end -= 1
    full_date = text[begin:end]
    date = full_date[0:10]
    return date


def get_user_birthdate(user_id):
    result = ""
    values = {
        "user_ids": user_id,
        "fields": "bdate"
    }
    response = VK_SESSION.method("users.get", values)
    if len(response) > 0:
        if "bdate" in response[0]:
            result = response[0]["bdate"]
        else:
            result = "No info about birthdate"
    else:
        result = "Error of receiving data"
    return result


def check_following(group_id, user_id):
    values = {
        "group_id": group_id,
        "user_id": user_id
    }
    response = VK_SESSION.method("groups.isMember", values)
    if response == 1:
        is_member = True
    else:
        is_member = False
    return is_member


def main():
    groups = { # вставить данные нужных сообществ
        "values": [
            {
                "id": 1,
                "name": "APIClub"
            }
        ]
    }

    profile_url = get_profile_url()
    screenname = select_screenname(profile_url)
    user_id = get_user_id(screenname)
    html_text = load_html(user_id)

    name = get_name(html_text)
    date = get_acc_date_created(html_text)
    birthdate = get_user_birthdate(user_id)
    print("Full name: " + name.decode("cp1251"))
    print("Account date created: " + str(date))
    print("User birthdate: " + str(birthdate))

    for group in groups["values"]:
        group_id = group["id"]
        group_name = group["name"]
        is_member = check_following(group_id, user_id)
        print("Is member " + group_name + ": " + str(is_member))


main()
