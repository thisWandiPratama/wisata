import React, { Component } from 'react';
import { View, Text, Image, TouchableOpacity, ScrollView, TextInput } from 'react-native';
import { removeToken } from '../serivce/token/token';
import { getService, postService} from "../serivce/apiTemplate"
import { baseURLPrimary } from "../serivce/ApiService"

class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      categori: [],
      search: false,
      name:""
    };
  }

  categori = () => {
    return this.state.categori.map((value, index) => {
      return (
        <TouchableOpacity onPress={() => this.props.navigation.navigate("ByCategory", { category_id: value.id, name_categori: value.name })} key={index} style={{ height: 120, width: 120, borderWidth: 2, borderRadius: 20, alignItems: "center", justifyContent: "center", margin: 10 }}>
          <Image
            style={{ height: 50, width: 80, borderRadius: 10 }}
            source={{ uri: `${baseURLPrimary}${value.avatar}` }}
          />
          <Text style={{ fontSize: 12, fontWeight: "bold", color: "#000", textAlign: "center" }}>{value.name.length < 10 ? value.name : value.name.slice(0, 50)}...</Text>
        </TouchableOpacity>
      )
    })
  }

  async componentDidMount() {
    const categori = await getService("all_category")
    console.log(categori)
    this.setState({ categori: categori.data })
  }

  search = async() => {
    if(this.state.name.length>0){
      this.props.navigation.navigate("Search",{name:this.state.name})
    }else{
      alert("Kolom search wajib diisi")
    } 
  }

  render() {
    return (
      <View style={{ flex: 1, }}>
        <View style={{ alignItems: 'center' }}>
          <View style={{ height: 80, width: "100%", alignItems: "center", justifyContent: 'space-between', flexDirection: "row" }}>
            <TouchableOpacity onPress={() => this.props.navigation.openDrawer()}>
              <Image style={{ height: 40, width: 40, marginLeft: 25 }} source={{ uri: "https://i.ibb.co/wMmB9cF/list.png" }} />
            </TouchableOpacity>
            {this.state.search == false ?
              <TouchableOpacity onPress={() => this.setState({ search: !this.state.search })}>
                <Image style={{ height: 40, width: 40, marginRight: 25 }} source={{ uri: "https://i.ibb.co/dMGhnQk/loupe.png" }} />
              </TouchableOpacity>
              :
              <View style={{ flexDirection: "row", height: 50, width: "80%", justifyContent:"space-around",borderBottomWidth:1}}>
                <TextInput
                  placeholder='Search'
                  style={{
                    height: 50,
                    width: "80%",
                    marginRight: 10
                  }}
                  onChangeText={(name) => this.setState({name:name})}
                />
                <TouchableOpacity onPress={() => this.search()}  >
                  <Image style={{ height: 40, width: 40, marginRight: 5 }} source={{ uri: "https://i.ibb.co/dMGhnQk/loupe.png" }} />
                </TouchableOpacity>
              </View>
            }
          </View>
        </View>
        <View style={{ flex: 1, }}>
          <View style={{ width: "100%", height: 100, borderBottomColor: "#aeaeae", borderBottomWidth: 2 }}>
            <View style={{
              shadowOffset: {
                width: 0,
                height: 2,
              },
              shadowOpacity: 0.25,
              shadowRadius: 3.84,

              elevation: 5,
              borderRadius: 10,
              height: 150,
              width: 150,
              marginLeft: 20,
              position: 'absolute',
              bottom: -80,
              shadowColor: "#000",
              alignItems: "center",
              justifyContent: "center"
            }}>
              <Image style={{
                borderRadius: 10,
                height: 150, width: 150, resizeMode: "contain",
              }} source={{ uri: "https://i.ibb.co/XD3t2t6/destination-1.png" }} />
            </View>
          </View>
          <View style={{ width: "100%", height: 50, alignItems: 'flex-end', justifyContent: 'center', marginTop: 20 }}>
            <View style={{ height: 50, width: "50%", alignItems: "center", justifyContent: "center", borderWidth: 1 }}>
              <Text>Rencanakan perjalananmu</Text>
            </View>
          </View>
          <View style={{ marginTop: 30, marginLeft: 30 }}>
            <Text onPress={() => {
              removeToken()
              this.props.navigation.replace("Login")
            }} style={{ fontSize: 25, color: "#000", fontWeight: "bold" }}>Daftar Kategori</Text>
          </View>
          <ScrollView>
            <View style={{ flex: 1, flexDirection: "row", flexWrap: "wrap", margin: 20, justifyContent: "center" }}>
              {this.categori()}
            </View>
          </ScrollView>
        </View>
      </View>
    );
  }
}

export default Home;
