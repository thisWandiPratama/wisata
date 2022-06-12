import React, { Component } from 'react';
import { View, Text, Image, TouchableOpacity, ToastAndroid } from 'react-native';
import { getWithTokenService } from '../serivce/apiTemplate'
import { getToken } from '../serivce/token/token'
import { baseURLPrimary } from '../serivce/ApiService'
import { choosePhotoFromLibrary } from '../serivce/helper';
import { uploadAvatarService } from '../serivce/ApiAuthService';

class Profile extends Component {
  constructor(props) {
    super(props);
    this.state = {
      token: "",
      profile: {}
    };
  }

  async componentDidMount() {
    const token = await getToken()
    const data = await getWithTokenService("profile", token)
    this.setState({
      profile: data.data
    })
  }

  uploadAvatar = async () => {
    const token = await getToken();
    const photo = await choosePhotoFromLibrary();
    const upload = await uploadAvatarService(token, photo.path, photo.mime);
    console.log(upload)
    ToastAndroid.show(upload.meta.message, ToastAndroid.SHORT);
    this.getProfile();
    return false;
  }


  getProfile = async () => {
    const token = await getToken()
    const data = await getWithTokenService("profile", token)
    this.setState({
      profile: data.data
    })
  }

  alertLogout = () => {
    Alert.alert(
      'Peringatan',
      'Apakah Anda Yakin Ingin Logout ?',
      [
        {
          text: 'Batal',
          onPress: () => ToastAndroid.show("Logout Cancled", ToastAndroid.SHORT),
        },
        {
          text: 'Logout', onPress: () => {
            logoutService();
            this.props.navigation.replace("Guest")
          }
        },
      ],
      { cancelable: false },
    );
  }

  render() {
    console.log(this.state.profile)
    return (
      <View style={{ flex: 1 }}>
        <TouchableOpacity onPress={() => this.uploadAvatar()} style={{ width: "100%", height: 250, alignItems: "center", justifyContent: "center", marginTop: 20 }}>
          <Image
            source={{
              uri: `${baseURLPrimary}${this.state.profile.avatar}`
            }}
            style={{
              height: 200,
              width: 200,
              borderRadius: 100
            }}
          />
          <Text style={{ fontSize: 20, marginTop: 30, fontWeight: "bold" }}> {this.state.profile.name} </Text>
        </TouchableOpacity>
        <View style={{ flex: 1, alignItems: "center" }}>
          <View style={{ height: 50, width: "80%", borderWidth: 1, justifyContent: "center", borderRadius: 10, marginTop: 20 }}>
            <Text style={{ fontSize: 20, fontWeight: "bold", marginLeft: 10, }}>{this.state.profile.email}</Text>
          </View>
          <View style={{ height: 50, width: "80%", borderWidth: 1, justifyContent: "center", borderRadius: 10, marginTop: 20 }}>
            <Text style={{ fontSize: 20, fontWeight: "bold", marginLeft: 10, }}>{this.state.profile.phone}</Text>
          </View>
        </View>
      </View>
    );
  }
}

export default Profile;
