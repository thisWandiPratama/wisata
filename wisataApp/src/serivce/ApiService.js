const baseURL = "http://192.168.1.9:8080/api/v1";
const baseURLPrimary = "http://192.168.1.9:8080/";

let myHeadersApiPublic = new Headers();
myHeadersApiPublic.append("Accept", "application/json");

export { baseURL, myHeadersApiPublic,baseURLPrimary };