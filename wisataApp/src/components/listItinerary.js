import React, { Component } from 'react';
import { View, Text, TouchableOpacity, Image, ScrollView } from 'react-native';
import { getService, postService, } from "../serivce/apiTemplate"
import {getUser} from '../serivce/token/token'


class ListItinerary extends Component {
  constructor(props) {
    super(props);
    this.state = {
      itinerary: []
    };
  }

  async componentDidMount() {
    const user_id = await getUser()
    const datasend = {
      'user_id':parseInt(user_id)
    }
    const data = await postService(datasend,"all_itinerary_by_user")
    // console.log(data.data)
    this.setState({
      itinerary: data.data
    })
  }

  getItinerary = async() => {
    const user_id = await getUser()
    const datasend = {
      'user_id':parseInt(user_id)
    }
    const data = await postService(datasend,"all_itinerary_by_user")
    this.setState({
      itinerary: data.data
    })
  }

  delete = async(id) => {
  const datasend = {
    "itinerary_id":id
  }
    const data = await postService(datasend,"delete_itinerary")
    console.log(data)
    this.getItinerary()
  }


  renderItinerary = () => {
    return this.state.itinerary.map((value, index) => {
      return (
        <TouchableOpacity key={index} disabled={true} style={{
          height: 250,
          width: "100%",
          borderWidth: 1,
          marginBottom: 10,
          borderRadius: 20,
        }}>
          <TouchableOpacity onPress={() => this.props.navigation.navigate("DetailItinerary",{itinerary:value.id})}>
            <View style={{
              height: 180,
              width: "100%",
            }}>
              <Image
                source={{
                  uri: `https://i.ibb.co/bPfz7zT/43715.jpg`
                }}
                style={{
                  height: "100%",
                  width: "100%",
                  resizeMode: "stretch",
                  borderRadius: 20
                }}
              />
            </View>
            <Text style={{ fontSize: 18, fontWeight: "bold", marginLeft: 10, color: "#000" }}>{value.intial_local}</Text>
          </TouchableOpacity>
          <View style={{ flexDirection: "row", width: "100%", height: 40, justifyContent: "flex-end" }}>
            <TouchableOpacity onPress={() => this.delete(value.id)}>
              <Image
                source={{
                  uri: `https://i.ibb.co/xjBfGX4/delete.png`
                }}
                style={{
                  height: 35,
                  width: 35,
                  resizeMode: "contain",
                  marginRight: 20
                }}
              />
            </TouchableOpacity>
            <Image
              source={{
                uri: `https://i.ibb.co/nPjYTHn/arrow-point-to-right.png`
              }}
              style={{
                height: 35,
                width: 35,
                resizeMode: "contain",
                marginRight: 30
              }}
            />
          </View>
        </TouchableOpacity>
      )
    })
  }


  render() {
    return (
      <View style={{ flex: 1 }}>
        <View style={{ height: 50, width: "100%", flexDirection: "row", marginTop: 10, alignItems: "center" }}>
          <TouchableOpacity onPress={() => this.props.navigation.goBack()} style={{ height: 50, width: 80, backgroundColor: "#aeaeae", alignItems: "center", justifyContent: "center", marginLeft: 10, borderRadius: 20 }}>
            <Text style={{ fontSize: 15, fontWeight: "bold" }}>Back</Text>
          </TouchableOpacity>
          <Text style={{ marginLeft: 50, fontSize: 20, fontWeight: "bold", }}> listItinerary </Text>
        </View>
        <ScrollView>
          <View style={{ flex: 1, flexDirection: "row", flexWrap: "wrap", margin: 20, justifyContent: "center" }}>
            {this.renderItinerary()}
          </View>
        </ScrollView>
      </View>
    );
  }
}

export default ListItinerary;
