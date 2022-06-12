import React, { Component } from 'react';
import { View, Text, TouchableOpacity, Image, TextInput, Modal, ScrollView, ActivityIndicator } from 'react-native';
import DatePicker from 'react-native-date-picker'

import { postService, getService } from '../serivce/apiTemplate'
import { getUser } from '../serivce/token/token'


class AddItinerary extends Component {
  constructor(props) {
    super(props);
    this.state = {
      initial_local: "",
      start_day: new Date(),
      end_day: new Date(),
      start_time: new Date(),
      end_time: new Date(),
      itinerary: [],
      user_id: "",
      // modal
      openstartday: false,
      openendday: false,
      openstarttime: false,
      openendtime: false,
      modal: false,
      modalDrawer: false,
      // default data
      category: [],
      allwisata: [],
      route: 0,
      time: new Date(),
      timedate: false,
      // objet
      iditinerary: 0,
      isloading: false,

      // timeline
      time1: "",
      title: "",
      description: "",
      latitude: "",
      longitude: "",
      isloading: true
    };
  }


  async componentDidMount() {
    const data = {
      "category_id": this.props.route.params.category_id
    }
    const allTouristByCategori = await postService(data, "all_tourist_by_categori")
    this.setState({ allwisata: allTouristByCategori.data })
    const categori = await getService("all_category")
    this.setState({ category: categori.data })
    setTimeout(() => {
      this.setState({
        isloading: false
      })
    }, 5000);
  }

  reload = async (category_id) => {
    this.setState({isloading:true})
    const data = {
      "category_id": category_id
    }
    const allTouristByCategori = await postService(data, "all_tourist_by_categori")
    if(allTouristByCategori.data.length>0){
      this.setState({ allwisata: allTouristByCategori.data })
    }else{
      this.setState({ allwisata:[]})
    }
    setTimeout(() => {
      this.setState({isloading:false})
    }, 5000);
  }

  itinerary = () => {
    return this.state.itinerary.map((value, index) => {
      return (
        <View key={index} style={{
          height: 50, width: "100%", flexDirection: "row", justifyContent: "space-between", borderBottomWidth: 1,
        }}>
          <View>
            <Text style={{ color: "red", fontWeight: "bold", fontSize: 15 }}>Tourist</Text>
            <Text style={{ color: "#000", fontWeight: "bold", fontSize: 12 }}>{value.description}</Text>
          </View>
          <View>
            <Text style={{ color: "red", fontWeight: "bold", fontSize: 15 }}>Day</Text>
            <Text style={{ color: "#000", fontWeight: "bold", fontSize: 12 }}>{value.time}</Text>
          </View>
          <View>
            <Text style={{ color: "red", fontWeight: "bold", fontSize: 15 }}>Time</Text>
            <Text style={{ color: "#000", fontWeight: "bold", fontSize: 12 }}>{value.title}</Text>
          </View>
          <TouchableOpacity onPress={() => this.setState({
            itinerary: this.state.itinerary.filter(items => items.id !== value.id)
          })} style={{ height: 50, alignItems: "center", justifyContent: "center" }}>
            <Text style={{ color: "red", fontWeight: "bold", fontSize: 15 }}> Delete</Text>
          </TouchableOpacity>
        </View >
      )
    })
  }

  listwisata = () => {
    return this.state.allwisata.map((value, index) => {
      return (
        <TouchableOpacity disabled={true} key={index} style={{ height: 100, width: "90%", flexDirection: "row", justifyContent: "space-between", }}>
          <View>
            <Text style={{ fontSize: 15, fontWeight: "bold", color: "red" }}>Tourist</Text>
            <Text style={{ fontSize: 15, fontWeight: "bold", color: "#000" }}>{value.name.length < 22 ? value.name : value.name.slice(0, 18)}...</Text>
          </View>
          <View style={{ height: 80, width: 80, marginLeft: 20, alignItems: "center", alignItems: "center" }}>
            <Text style={{ fontSize: 15, fontWeight: "bold", color: "red" }}>Day</Text>
            <View style={{ flexDirection: "row" }}>
              <TouchableOpacity style={{ marginRight: 5, width: 20, height: 20, alignItems: "center" }} TouchableOpacity onPress={() => this.setState({ route: this.state.route - 1 })}>
                <Text>-</Text>
              </TouchableOpacity>
              <Text>{this.state.route}</Text>
              <TouchableOpacity style={{ marginLeft: 5, width: 20, height: 20, alignItems: "center" }} TouchableOpacity onPress={() => this.setState({ route: this.state.route + 1 })}>
                <Text>+</Text>
              </TouchableOpacity>
            </View>
          </View>
          <View style={{ marginRight: 10 }}>
            <Text style={{ fontSize: 15, fontWeight: "bold", color: "red" }}>Time</Text>
            <DatePicker
              modal
              mode={"time"}
              open={this.state.timedate}
              date={this.state.time}
              onConfirm={(date) => {
                this.setState({ timedate: !this.state.timedate, time: date })
              }}
              onCancel={() => {
                this.setState({ timedate: !this.state.timedate })
              }}
            />
            <TouchableOpacity onPress={() => this.setState({ timedate: !this.state.timedate })}>
              <Text>{this.state.time.getHours()}:{this.state.time.getMinutes()}</Text>
            </TouchableOpacity>
          </View>
          <View style={{ height: 50, justifyContent: "center" }}>
            <TouchableOpacity style={{
              marginLeft: 20
            }} onPress={() => {
              this.setState({
                time1: `Day ${this.state.time1}`,
                title: this.state.time.getHours() + ":" + this.state.time.getMinutes(),
                description: `${value.name}`,
                latitude: `${value.latitude}`,
                longitude: `${value.longitude}`
              })
              this.state.itinerary.push(

                { id: this.state.iditinerary + 1, description: `${value.name}`, time: `Day ${this.state.route}`, title: this.state.time.getHours() + ":" + this.state.time.getMinutes(), latitude: `${value.latitude}`, longitude: `${value.longitude}` }
              )
              this.setState({ modal: !this.state.modal, iditinerary: this.state.iditinerary + 1 })
            }}>
              <Text style={{ fontSize: 15, fontWeight: "bold", color: "red" }}>Add Itinerary</Text>
            </TouchableOpacity>
          </View>
        </TouchableOpacity>
      )
    })
  }

  listcategory = () => {
    return this.state.category.map((value, index) => {
      return (
        <TouchableOpacity onPress={() => this.reload(value.id)} key={index} style={{
          height: 50, width: 100, alignItems: "center", justifyContent: "center", marginRight: 10,
          shadowColor: "#000",
          shadowOffset: {
            width: 0,
            height: 3,
          },
          shadowOpacity: 0.29,
          shadowRadius: 4.65,

          elevation: 7,
          backgroundColor: "#aeaeae",
          borderRadius: 25
        }}>
          <Text style={{ fontSize: 10, textAlign: "center" }}>{value.name}</Text>
        </TouchableOpacity>
      )
    })
  }

  create = async () => {
    this.setState({ isloading: true })
    const user_id = await getUser()
    // console.log(this.state.initial_local, this.state.start_day, this.state.end_day, this.state.start_time, this.state.end_time, this.state.itinerary)
    if (this.state.initial_local.length == 0 || this.state.start_day.length == 0 || this.state.end_day.length == 0 || this.state.start_time.length == 0 || this.state.end_time.length == 0 || this.state.itinerary.length == 0) {
      this.setState({ isloading: false })
      alert("Semua data wajib terisi")
    } else {
      const datasend = {
        intial_local: this.state.initial_local,
        start_day: `${this.state.start_day.getDate()}/${this.state.start_day.getMonth()}/${this.state.start_day.getFullYear()}`,
        end_day: `${this.state.end_day.getDate()}/${this.state.end_day.getMonth()}/${this.state.end_day.getFullYear()}`,
        start_time: `${this.state.start_time.getHours()}:${this.state.start_time.getMinutes()}`,
        end_time: `${this.state.end_time.getHours()}: ${this.state.end_day.getMinutes()}`,
        user_id: parseInt(user_id)
      }
      const postItinerary = await postService(datasend, "add_itinerary")
      console.log(postItinerary)
      if (postItinerary.meta.code == 200 && postItinerary.data.id) {
        var timer = 0
        for (var i = 0; i < this.state.itinerary.length; i++) {
          timer += 1
          console.log("timer:" + timer + " " + "lenght:" + this.state.itinerary.length)
          const root = this.state.itinerary[i]
          const datasend = {
            "itinerary_id": postItinerary.data.id,
            "time": root.time,
            "title": root.title,
            "description": root.description,
            "latitude": root.latitude,
            "longitude": root.longitude
          }
          const timeline = await postService(datasend, "add_timeline")
          if (timer == this.state.itinerary.length) {
            setTimeout(() => {
              this.setState({ isloading: false })
              this.props.navigation.navigate("ListItinerary")
            }, 3000);
          }
        }

      } else {
        setTimeout(() => {
          this.setState({ isloading: false })
          alert(`${postItinerary.meta.message}`)
          this.setState({ isloading: false })
        }, 3000);
      }
    }
  }

  render() {
    console.log(this.state.itinerary)
    return (
      <View style={{ flex: 1, }}>
        <ScrollView>
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
          <View style={{ margin: 20 }}>
            <Text style={{ textAlign: "center", fontSize: 20, fontWeight: "bold", color: "#000" }}> Buat Itinerary </Text>
            <View style={{ height: 50, width: "100%", flexDirection: "row", alignItems: "center" }}>
              <Image
                source={{
                  uri: `https://i.ibb.co/Fxz89vd/placeholder.png`
                }}
                style={{
                  height: 30,
                  width: 30,
                  marginRight: 5
                }}
              />
              <TextInput
                placeholder='Lokasi Awal'
                style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10 }}
                onChangeText={(awal) => this.setState({ initial_local: awal })}
              />
            </View>
            <View style={{ flexDirection: "row", justifyContent: "space-between", marginTop: 10 }}>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center" }}>
                <Image
                  source={{
                    uri: `https://i.ibb.co/9v81cwT/schedule.png`
                  }}
                  style={{
                    height: 30,
                    width: 30,
                    marginRight: 5
                  }}
                />
                <DatePicker
                  modal
                  mode={"date"}
                  open={this.state.openstartday}
                  date={this.state.start_day}
                  onConfirm={(date) => {
                    this.setState({ openstartday: !this.state.openstartday, start_day: date })
                  }}
                  onCancel={() => {
                    this.setState({ openstartday: !this.state.openstartday })
                  }}
                />
                <TouchableOpacity onPress={() => this.setState({ openstartday: !this.state.openstartday })} style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}>
                  <Text>{this.state.start_day.getDate()}/{this.state.start_day.getMonth()}/{this.state.start_day.getFullYear()}</Text>
                </TouchableOpacity>
              </View>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center", marginRight: 25 }}>
                <Image
                  source={{
                    uri: `https://i.ibb.co/9v81cwT/schedule.png`
                  }}
                  style={{
                    height: 30,
                    width: 30,
                    marginRight: 5
                  }}
                />
                <DatePicker
                  modal
                  mode={"date"}
                  open={this.state.openendday}
                  date={this.state.end_day}
                  onConfirm={(date) => {
                    this.setState({ openendday: !this.state.openendday, end_day: date })
                  }}
                  onCancel={() => {
                    this.setState({ openendday: !this.state.openendday })
                  }}
                />
                <TouchableOpacity onPress={() => this.setState({ openendday: !this.state.openendday })} style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}>
                  <Text>{this.state.end_day.getDate()}/{this.state.end_day.getMonth()}/{this.state.end_day.getFullYear()}</Text>
                </TouchableOpacity>
              </View>
            </View>
            <View style={{ flexDirection: "row", justifyContent: "space-between", marginTop: 10 }}>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center" }}>
                <Image
                  source={{
                    uri: `https://i.ibb.co/mcVd6gL/start.png`
                  }}
                  style={{
                    height: 30,
                    width: 30,
                    marginRight: 5
                  }}
                />
                <DatePicker
                  modal
                  mode={"time"}
                  open={this.state.openstarttime}
                  date={this.state.start_time}
                  onConfirm={(date) => {
                    this.setState({ openstarttime: !this.state.openstarttime, start_time: date })
                  }}
                  onCancel={() => {
                    this.setState({ openstarttime: !this.state.openstarttime })
                  }}
                />
                <TouchableOpacity onPress={() => this.setState({ openstarttime: !this.state.openstarttime })} style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}>
                  <Text>{this.state.start_time.getHours()}:{this.state.start_time.getMinutes()}</Text>
                </TouchableOpacity>
              </View>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center", marginRight: 25 }}>
                <Image
                  source={{
                    uri: `https://i.ibb.co/mcVd6gL/start.png`
                  }}
                  style={{
                    height: 30,
                    width: 30,
                    marginRight: 5
                  }}
                />
                <DatePicker
                  modal
                  mode={"time"}
                  open={this.state.openendtime}
                  date={this.state.end_time}
                  onConfirm={(date) => {
                    this.setState({ openendtime: !this.state.openendtime, end_time: date })
                  }}
                  onCancel={() => {
                    this.setState({ openendtime: !this.state.openendtime })
                  }}
                />
                <TouchableOpacity onPress={() => this.setState({ openendtime: !this.state.openendtime })} style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}>
                  <Text>{this.state.end_time.getHours()}:{this.state.end_time.getMinutes()}</Text>
                </TouchableOpacity>
              </View>
            </View>
            <View style={{ height: 50, width: "100%", justifyContent: "space-between", flexDirection: "row" }}>
              <Text style={{ textAlign: "center", fontSize: 20, fontWeight: "bold", color: "#000", marginTop: 20 }}> Wisata </Text>
              <TouchableOpacity onPress={() => this.setState({ modal: !this.state.modal })}>
                <Text style={{ textAlign: "center", fontSize: 20, fontWeight: "bold", color: "#000", marginTop: 20 }}> Add </Text>
              </TouchableOpacity>
            </View>
            {this.state.itinerary.length == 0 ? <Text style={{ textAlign: "center", color: "#aeaeae" }}>Tambahkan daftar kunjungan</Text> :
              this.itinerary()
            }
          </View>
        </ScrollView>
        <View style={{ height: 50, width: "100%", position: "absolute", alignItems: "center", justifyContent: "center", bottom: 10, backgroundColor: "transparent" }}>
          <TouchableOpacity onPress={() => this.create()} style={{ height: 50, width: "70%", backgroundColor: "#aeaeae", borderRadius: 20, alignItems: "center", justifyContent: "center" }}>
            {this.state.isloading ? <ActivityIndicator /> : <Text style={{ fontSize: 20, fontWeight: "bold", color: "#000" }}>Create</Text>}
          </TouchableOpacity>
        </View>
        <Modal
          animationType="slide"
          visible={this.state.modal}
          onRequestClose={() => {
            this.setState({ modal: !this.state.modal });
          }}
        >
          <View style={{ flex: 1 }}>
            <Text style={{ fontSize: 20, textAlign: "center", marginTop: 20, fontWeight: "bold", color: "#000" }}>Pilih Tujuan Kunjungan</Text>
            <View style={{ margin: 20, height: 50, flexDirection: "row", width: "100%" }}>
              <ScrollView horizontal={true}>
                {this.listcategory()}
              </ScrollView>
            </View>
            <View style={{ flex: 1, margin: 20, marginTop: 0, }}>
              <ScrollView>
                {this.state.isloading ? <View style={{ flex: 1, alignItems: 'center', justifyContent: "center" }}>
                  <ActivityIndicator size="large" />
                </View> : this.state.allwisata.length>0?this.listwisata():<Text style={{fontSize:20, textAlign:"center"}}>Tidak ada data</Text>}
              </ScrollView>
            </View>
          </View>
        </Modal>

      </View>
    );
  }
}

export default AddItinerary;
