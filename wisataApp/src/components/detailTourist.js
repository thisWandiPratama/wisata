import React, { Component } from 'react';
import { View, Text, TouchableOpacity, Image, ScrollView, Linking } from 'react-native';
import { postService } from "../serivce/apiTemplate"
import { baseURLPrimary } from "../serivce/ApiService"
import { WebView } from 'react-native-webview';


class DetailTourist extends Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      name: "",
      description: "",
      address: "",
      email: "",
      phone: "",
      website: "",
      latitude: "",
      longitude: "",
      link_video: "",
      image_primary: "",
      image_gallery: []
    };
  }

  async componentDidMount() {
    this.setState({
      id: this.props.route.params.id,
      name: this.props.route.params.name,
      description: this.props.route.params.description,
      address: this.props.route.params.address,
      email: this.props.route.params.email,
      phone: this.props.route.params.phone,
      website: this.props.route.params.website,
      latitude: this.props.route.params.latitude,
      longitude: this.props.route.params.longitude,
      link_video: this.props.route.params.link_video,
      image_primary: this.props.route.params.image_primary
    })

    const data = {
      "tourist_id": parseInt(this.state.id)
    }

    const image_gallery = await postService(data, "all_image_gallery_tourist")
    this.setState({ image_gallery: image_gallery.data })
  }

  image_gallery = () => {
    return this.state.image_gallery.map((value, index) => {
      return (
        <Image
          key={index}
          source={{
            uri: `${baseURLPrimary}${value.avatar}`
          }}
          style={{
            height: 250,
            width: "100%",
          }}
        />
      )
    })
  }

  video = () => {
    return <WebView source={{ uri: `${this.state.link_video}` }} />;
  }

  render() {
    return (
      <View style={{
        flex: 1,
      }}>
        <TouchableOpacity style={{
          height: 50,
          width: 70,
          margin: 20,
          backgroundColor: "#aeaeae",
          justifyContent: "center",
          alignItems: "center",
          borderRadius: 25
        }}
          onPress={() => this.props.navigation.goBack()}
        >
          <Text style={{
            fontSize: 16,
            color: "#000",
            fontWeight: "bold"
          }}>Back</Text>
        </TouchableOpacity>
        <View style={{ width: "100%", height: 250, justifyContent: "center" }}>
          <Image
            source={{
              uri: `${baseURLPrimary}${this.state.image_primary}`
            }}
            style={{
              height: 250,
              width: "100%",
            }}
          />
        </View>
        <View style={{ flex: 1, marginTop: 30, }}>
          <ScrollView>
            <View style={{ margin: 20 }}>
              <View style={{ width: "100%", justifyContent: "space-between", flexDirection: "row" }}>
                <View style={{ width: "90%" }}>
                  <Text style={{ fontSize: 20, fontWeight: "bold", color: "#000" }}>{this.state.name}</Text>
                </View>
                <TouchableOpacity
                  onPress={() => Linking.openURL(`${this.state.website
                    }`)}
                >
                  <Image
                    source={{
                      uri: `https://i.ibb.co/4KHcVkV/web-link.png`
                    }}
                    style={{
                      height: 30,
                      width: 30,
                    }}
                  />
                </TouchableOpacity>
              </View>
              <Text style={{ fontSize: 10, color: "#aeaeae", textAlign: "justify" }}>{this.state.address}.Phone:{this.state.phone},Email:{this.state.email}</Text>
              <Text style={{ fontSize: 15, color: "#000", textAlign: "justify" }}>{this.state.description}. </Text>
              {this.state.link_video.length == 0 ? null :
                <View style={{ width: "100%", height: 500 }}>
                  {this.video()}
                </View>
              }
              {this.image_gallery()}

            </View>

          </ScrollView>
          <View style={{height:50, width:"100%",position:"absolute", backgroundColor:"transparent",alignItems:"center", bottom:10, justifyContent:"center"}}>
          <TouchableOpacity onPress={()=> Linking.openURL(`https://www.google.com/maps/search/?api=1&query=${this.state.latitude},${this.state.longitude}`)} style={{ height: 50, width: "50%", backgroundColor: "#aeaeae", alignItems:"center", justifyContent:"center", borderRadius:20}}>
            <Text style={{fontSize:20, color:"red", fontWeight:"bold"}}>Direction</Text>
          </TouchableOpacity>
          </View>
        </View>
      </View>
    );
  }
}

export default DetailTourist;
