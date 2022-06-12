import React, { Component } from 'react';
import { View, Text, ScrollView, TouchableOpacity, Image } from 'react-native';
import { postService } from '../serivce/apiTemplate'
import { baseURLPrimary } from "../serivce/ApiService"

class Bycategory extends Component {
  constructor(props) {
    super(props);
    this.state = {
      category_id: 0,
      name_categori: "",
      alltourist: []
    };
  }

  async componentDidMount() {
    this.setState({
      category_id: this.props.route.params.category_id,
      name_categori: this.props.route.params.name_categori
    })
    const data = {
      "category_id": this.props.route.params.category_id
    }
    const allTouristByCategori = await postService(data, "all_tourist_by_categori")
    this.setState({ alltourist: allTouristByCategori.data })
  }

  allTouristByCategori = () => {
    return this.state.alltourist.map((value, index) => {
      return (
        <TouchableOpacity key={index} style={{
          width: "90%", height: 300, borderWidth: 2, borderRadius: 10, shadowOffset: {
            width: 0,
            height: 2,
          },
          shadowOpacity: 0.25,
          shadowRadius: 3.84,

          elevation: 5,
          alignItems: "center"
        }}

          onPress={() => this.props.navigation.navigate("DetailTourist", {
            id: value.category_id,
            name: value.name,
            description: value.description,
            address: value.address,
            email: value.email,
            phone: value.phone,
            website: value.website,
            latitude: value.latitude,
            longitude: value.longitude,
            link_video: value.link_video,
            image_primary: value.image_primary
          })}

        >
          <View style={{ height: 200, width: "98%", borderWidth: 2, marginTop: 5, borderRadius: 7, alignItems: "center", justifyContent: "center" }}>
            <Image
              source={{ uri: `${baseURLPrimary}${value.image_primary}` }}
              style={{ height: 190, width: "98%", borderRadius: 10 }}
            />
          </View>
          <View style={{ height: 85, width: "98%", }}>
            <View style={{ height: 43, width: "98%" }}>
              <Text style={{ color: "#000", fontSize: 20, fontWeight: "bold" }}>{value.name}</Text>
              <Text style={{ color: "#aeaeae", fontSize: 10, fontWeight: "bold" }}>{value.address.slice(0, 140)}...</Text>
            </View>
            <View style={{ height: 43, width: "98%", alignItems: "flex-end", justifyContent: "flex-end" }}>
              <Image
                source={{ uri: `https://i.ibb.co/wgYbXhF/right-arrow.png` }}
                style={{ height: 25, width: 25, resizeMode: "contain", marginBottom: 5 }}
              />
            </View>
          </View>
        </TouchableOpacity>
      )
    })
  }

  render() {
    return (
      <View style={{ flex: 1 }}>
        <View style={{ height: 100, width: "100%", flexDirection: "row", justifyContent: "space-between", alignItems: "center" }}>
          <View style={{ width: "80%" }}>
            <Text style={{
              fontSize: 30, color: "#000", marginTop: 20, marginLeft: 20, fontWeight: "bold", shadowOffset: {
                width: 0,
                height: 2,
              },
              shadowOpacity: 0.25,
              shadowRadius: 3.84,

              elevation: 5,
            }}> {this.state.name_categori} </Text>
          </View>
          <TouchableOpacity onPress={() => this.props.navigation.navigate("AddItinerary",{category_id:this.state.category_id})}>
            <Image
              source={{ uri: `https://i.ibb.co/ySzZDzq/bookmark.png` }}
              style={{ height: 50, width: 50, marginRight: 20 }}
            />
          </TouchableOpacity>
        </View>
        <ScrollView>
          <View style={{ flex: 1, flexDirection: "row", flexWrap: "wrap", margin: 20, justifyContent: "center" }}>
            {this.allTouristByCategori()}
          </View>
        </ScrollView>
      </View>
    );
  }
}

export default Bycategory;
