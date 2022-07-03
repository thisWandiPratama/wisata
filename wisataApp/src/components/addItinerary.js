import React, { Component } from 'react';
import { View, Text, TouchableOpacity, Image, TextInput, Modal, ScrollView, ActivityIndicator, Alert } from 'react-native';
import DatePicker from 'react-native-date-picker'
import Geolocation from '@react-native-community/geolocation';

import { postService, getService } from '../serivce/apiTemplate'
import { getUser } from '../serivce/token/token'
import { getDistance, getPreciseDistance, getSpeed } from 'geolib';

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
      // itinerary: [{ "description": "Pantai Hamadi", "id": 1, "jarak": 15657791, "latitude": "2.5264303", "longitude": "140.6120972", "time": "Day 0", "title": "20:56" }, { "description": "Puncak Harfat", "id": 2, "jarak": 14493458, "latitude": "-1.97346", "longitude": "130.465494", "time": "Day 0", "title": "20:56" }, { "description": "Pantai Harlem", "id": 3, "jarak": 15579576, "latitude": "-2.4620136", "longitude": "140.3677778", "time": "Day 0", "title": "20:56" }],
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
      jarakTempuh: "",
      isloading: true,
      awal_latitude: "",
      awal_longitude: ""
    };
  }


  async componentDidMount() {
    Geolocation.getCurrentPosition(
      position => {
        const initialPosition = JSON.stringify(position);
        this.setState({
          awal_latitude: `${JSON.parse(initialPosition).coords.latitude}`,
          awal_longitude: `${JSON.parse(initialPosition).coords.longitude}`
        });
      },
      error => Alert.alert('Error', "Gagal mendapatkan koordinat lokasi Anda saat ini. Silakan input manual koordinat lokasi titik awal"),
      { enableHighAccuracy: true, timeout: 20000, },
    );
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
    this.setState({ isloading: true })
    const data = {
      "category_id": category_id
    }
    const allTouristByCategori = await postService(data, "all_tourist_by_categori")
    if (allTouristByCategori.data.length > 0) {
      this.setState({ allwisata: allTouristByCategori.data })
    } else {
      this.setState({ allwisata: [] })
    }
    setTimeout(() => {
      this.setState({ isloading: false })
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
            <Text style={{ fontSize: 15, fontWeight: "bold", color: "#000" }}>{value.name.length > 50 ? value.name.slice(0, 50) + "..." : value.name}</Text>
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
                { id: this.state.iditinerary + 1, description: `${value.name}`, time: `Day ${this.state.route}`, title: this.state.time.getHours() + ":" + this.state.time.getMinutes(), latitude: `${value.latitude}`, longitude: `${value.longitude}`, jarak: this.calculatePreciseDistance(value.latitude, value.longitude) }
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
    this.state.itinerary
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
          const timeMenit = this.calculatePreciseDistance(root.latitude, root.longitude).toFixed()
          const toString =String(timeMenit).slice(0,5)
          console.log("time menit "+toString)
          const toJam = toString/60
          console.log("jam "+toJam.toFixed(1))

          var fixTime = ""

          if (parseInt(timeMenit <60)){
            fixTime = `${toString} minute`
          }else{
            fixTime = `${toJam.toFixed(1)} hours`
          }

          const datasend = {
            "itinerary_id": postItinerary.data.id,
            "time": root.time,
            // "time": root.time,
            "title": fixTime,
            "description": root.description,
            "latitude": root.latitude,
            "longitude": root.longitude,
            "jarak": this.calculatePreciseDistance(root.latitude, root.longitude)
          }

          const timeline = await postService(datasend, "add_timeline")
          if (timer == this.state.itinerary.length) {
            setTimeout(() => {
              this.setState({ isloading: false })
              this.props.navigation.navigate("ListItinerary")
            }, 10000);
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

  calculatePreciseDistance = (latitudeTourist, longitudeTourist) => {
    console.log(this.state.awal_latitude, this.state.awal_longitude)
    console.log(latitudeTourist, longitudeTourist)
    var pdis = getPreciseDistance(
      { latitude: this.state.awal_latitude, longitude: this.state.longitude },
      { latitude: latitudeTourist, longitude: longitudeTourist },
    );

    return pdis
  };

  calculateDuration = (latitudeTourist, longitudeTourist) => {
    var time = getSpeed(
      // { latitude: this.state.awal_latitude, longitude: this.state.longitude, time: 1360231200880 },
      // { latitudeTourist, longitude: longitudeTourist, time: 1360245600880 }
      { latitude: 51.567294, longitude: 7.38896, time: 1360231200880 },
      { latitude: 52.54944, longitude: 13.468509, time: 1360245600880 }

    );
    return time
  }

  // test = () => {
  //   var jarak = []
  //   this.state.itinerary.map((value, index) => {
  //     // jarak.push(value.jarak)
  //     // return jarak
  //     console.log(value.jarak)
  //   })
  // }

  render() {
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
                placeholder='Nama Lokasi Awal'
                style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10 }}
                onChangeText={(awal) => this.setState({ initial_local: awal })}
              />
            </View>
            <View style={{ flexDirection: "row", justifyContent: "space-between", marginTop: 10 }}>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center" }}>
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
                <View
                  style={{
                    width: "100%",
                    marginBottom: 5
                  }}
                >
                  <Text
                    style={{
                      fontSize: 10
                    }}
                  >Latitude titik awal</Text>
                  <TextInput
                    value={this.state.awal_latitude}
                    placeholder='Latitude Awal'
                    style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}
                    onChangeText={(awal_latitude) => this.setState({ awal_latitude: awal_latitude })}
                  />
                </View>
              </View>
              <View style={{ height: 50, width: "40%", flexDirection: "row", alignItems: "center", marginRight: 25 }}>
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
                <View
                  style={{
                    width: "100%",
                    marginBottom: 5
                  }}
                >
                  <Text style={{
                    fontSize: 10
                  }}>Longitude titik awal</Text>
                  <TextInput
                    value={this.state.awal_longitude}
                    placeholder='Longitude Awal'
                    style={{ height: 50, width: "90%", borderWidth: 1, borderRadius: 10, justifyContent: "center", alignItems: "center" }}
                    onChangeText={(awal_longitude) => this.setState({ awal_longitude: awal_longitude })}
                  />
                </View>
              </View>
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
            <TouchableOpacity onPress={() => this.calculatePreciseDistance()}>
              <Text>Test</Text>
            </TouchableOpacity>
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
                </View> : this.state.allwisata.length > 0 ? this.listwisata() : <Text style={{ fontSize: 20, textAlign: "center" }}>Tidak ada data</Text>}
              </ScrollView>
            </View>
          </View>
        </Modal>

      </View>
    );
  }
}

export default AddItinerary;
