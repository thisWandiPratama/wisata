import AsyncStorage from "@react-native-async-storage/async-storage";

const storeToken = async (token) => {
    await AsyncStorage.setItem("token",token);
};

const storeid = async (token) => {
    await AsyncStorage.setItem("userid",JSON.stringify(token));
};

const getToken = async () => {
    const value = await AsyncStorage.getItem("token");
    return value;
};


const getUser = async () => {
    const value = await AsyncStorage.getItem("userid");
    return value;
};

const removeToken = async () => {
    await AsyncStorage.removeItem("token");
    await AsyncStorage.removeItem("userid");
};

export { storeToken, getToken, removeToken,storeid,getUser};