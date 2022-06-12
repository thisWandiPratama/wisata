import { myHeadersApiPublic, baseURL } from "./ApiService";
import { removeToken, storeToken } from "./token/token";
import { ToastAndroid } from "react-native";

// Logout Service
export const logoutService = () => {
    removeToken();
};


// Register Service
export const registerService = (data) => {
    const formData = data;
    const consume = fetch(`${baseURL}/register`, {
        method: "POST",
        body: JSON.stringify(formData),
        headers: myHeadersApiPublic,
    })
    .then(response => response.json())
    .catch(err =>  ToastAndroid.show("Network Request Failed!", ToastAndroid.SHORT));
    return consume;
};


// Login Service
export const loginService = (data) => {
    const formData = data;
    const consume = fetch(`${baseURL}/login`, {
        method: "POST",
        body: JSON.stringify(formData),
        headers: myHeadersApiPublic,
    })
    .then(response => response.json())
    .catch(err => ToastAndroid.show("Network Request Failed!", ToastAndroid.SHORT));
    return consume;
}


export const uploadAvatarService = (token, uri, type) => {
    let myHeadersApiPrivate = new Headers();
    myHeadersApiPrivate.append("Accept", "multipart/form-data");
    myHeadersApiPrivate.append("Authorization", `Bearer ${token}`);

    let formData = new FormData();

    formData.append("avatar", {
        name: uri,
        uri: uri,
        type: type,
    });

    const consume = fetch(`${baseURL}/updateavatar`, {
        method: "POST",
        body: formData,
        headers: myHeadersApiPrivate,
    })
    .then(response => response.json())
    .catch(err => console.log(err))
    return consume;
}