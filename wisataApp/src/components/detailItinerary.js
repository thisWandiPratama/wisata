import React, { Component } from 'react';
import { View, Text, TouchableOpacity, ScrollView,Linking } from 'react-native';
import Timeline from 'react-native-timeline-flatlist'
import { postService } from '../serivce/apiTemplate'
import { WebView } from 'react-native-webview';

class DetailItinerary extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      tab: 1
    };
  }

  async componentDidMount() {
    const datasend = {
      "itinerary_id": this.props.route.params.itinerary
    }
    const timeline = await postService(datasend, "timeline_by_itinerary_id")
    this.setState({
      data: timeline.data
    })

  }

  viewmap = (latitude, longitude) => {
    return <WebView source={{ uri: `https://www.google.com/maps/search/?api=1&query=${latitude},${longitude}` }} />;
  }

  maps = () => {
    return this.state.data.map((value, index) => {
      return (
        <TouchableOpacity style={{ height: 600, width: "100%", marginTop: 10 }}>
          <Text style={{ fontSize: 25, fontWeight: "bold", color: "#000" }}>{value.time}</Text>
          <View style={{ width: "100%", height: 500 }}>
            {this.viewmap(value.latitude, value.longitude)}
          </View>
          <View style={{alignItems:'center'}}>
          <TouchableOpacity 
          onPress={() => Linking.openURL(`https://www.google.com/maps/search/?api=1&query=${value.latitude},${value.longitude}`)}
          style={{
            height:50, width:"80%",
            backgroundColor:"#aeaeae",
            borderRadius:20,
            alignItems:"center",
            justifyContent:"center"
          }}>
            <Text style={{fontSize:25, fontWeight:"bold",color:"red"}}>Direction {value.time}</Text>
          </TouchableOpacity>
          </View>
        </TouchableOpacity>
      )
    })
  }


  render() {
    console.log(this.state.data)
    return (
      <View style={{ flex: 1 }}>
        <View style={{ height: 50, width: "100%", flexDirection: "row", marginTop: 10, alignItems: "center" }}>
          <TouchableOpacity onPress={() => this.props.navigation.goBack()} style={{ height: 50, width: 80, backgroundColor: "#aeaeae", alignItems: "center", justifyContent: "center", marginLeft: 10, borderRadius: 20 }}>
            <Text style={{ fontSize: 15, fontWeight: "bold" }}>Back</Text>
          </TouchableOpacity>
        </View>
        <View style={{ flex: 1 }}>
          <View style={{ height: 50, width: "100%", borderBottomWidth: 2, borderTopWidth: 2, marginTop: 50, borderColor: "#000", flexDirection: "row", alignItems: "center", justifyContent: "space-around" }}>
            <TouchableOpacity onPress={() => this.setState({ tab: 1 })} style={{ width: "30%", alignItems: "center", justifyContent: "center", height: 40, borderWidth: 1, borderRadius: 20 }}>
              <Text style={{ fontSize: 20, fontWeight: "bold", color: "#000" }}>Itinerary</Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={() => this.setState({ tab: 2 })} style={{ width: "30%", alignItems: "center", justifyContent: "center", height: 40, borderWidth: 1, borderRadius: 20 }}>
              <Text style={{ fontSize: 20, fontWeight: "bold", color: "#000" }}>Map</Text>
            </TouchableOpacity>
          </View>
          {this.state.tab == 1 ?
            <View style={{ marginTop: 20, marginLeft: 50 }}>
              <ScrollView>
                <Timeline
                  data={this.state.data}
                  circleSize={20}
                  circleColor='rgb(45,156,219)'
                  lineColor='rgb(45,156,219)'
                  timeContainerStyle={{ minWidth: 52, marginTop: -5 }}
                  timeStyle={{ textAlign: 'center', backgroundColor: '#ff9797', color: 'white', padding: 5, borderRadius: 13 }}
                  descriptionStyle={{ color: 'gray' }}
                  options={{
                    style: { paddingTop: 5 }
                  }}
                  isUsingFlatlist={true}
                />
              </ScrollView>
            </View>
            :
            <ScrollView>
              <View style={{ marginTop: 20, alignItems: "center" }}>
                {this.maps()}
              </View>
            </ScrollView>
          }
        </View>
      </View>
    );
  }
}

export default DetailItinerary;
