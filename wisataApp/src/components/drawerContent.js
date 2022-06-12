import React, { Component } from 'react';
import { View, Text, Image, TouchableOpacity } from 'react-native';
import { getWithTokenService, getDataToken,} from "../serivce/apiTemplate"
import { baseURLPrimary } from '../serivce/ApiService'
import {removeToken} from '../serivce/token/token'

class DrawerContent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      profile: {}
    };
  }

  async componentDidMount() {
    const tokendata = await getDataToken()
    if (tokendata.length > 0) {
      const data = await getWithTokenService("profile", tokendata)
      if (data.meta.status = "success") {
        this.setState({ profile: data.data })
      }
    }
  }


  render() {
    return (
      <View style={{ flex: 1 }}>
        <View style={{width:"100%", height:200, alignItems:"center", justifyContent:"center"}}>
          <Image
            source={{
              uri: `${baseURLPrimary}${this.state.profile.avatar}`
            }}
            style={{
              height: 100,
              width: 100
            }}
          />
          <Text style={{fontSize:20, marginTop:30, fontWeight:"bold"}}> {this.state.profile.name} </Text>
        </View>
        <View style={{flex:1,alignItems:"center"}}>
            <TouchableOpacity onPress={()=> this.props.navigation.navigate("Homedata")} style={{height:50, width:"80%", borderWidth:1, justifyContent:"center",borderRadius:10,marginTop:20}}>
              <Text style={{fontSize:20, fontWeight:"bold", marginLeft:10,}}>Home</Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={()=> this.props.navigation.navigate("ListItinerary")} style={{height:50, width:"80%", borderWidth:1, justifyContent:"center",borderRadius:10,marginTop:20}}>
              <Text style={{fontSize:20, fontWeight:"bold", marginLeft:10,}}>Itinerary</Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={()=> this.props.navigation.navigate("Profile")} style={{height:50, width:"80%", borderWidth:1, justifyContent:"center",borderRadius:10,marginTop:20}}>
              <Text style={{fontSize:20, fontWeight:"bold", marginLeft:10,}}>Profile</Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={()=> {
              removeToken()
              this.props.navigation.replace("Login")
            }} style={{height:50, width:"80%", borderWidth:1, justifyContent:"center",borderRadius:10,marginTop:20}}>
              <Text style={{fontSize:20, fontWeight:"bold", marginLeft:10,}}>Logout</Text>
            </TouchableOpacity>
        </View>
      </View>
    );
  }
}

export default DrawerContent;
