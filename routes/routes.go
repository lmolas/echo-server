package routes

import (
	"encoding/json"
	"net/http"

	"github.com/swaggo/swag"

	"github.com/labstack/echo/v4"
)

var demoBody = json.RawMessage([]byte(`[
  {
    "_id": "5fd1fa7e5a29916f09dcb788",
    "index": 0,
    "guid": "c321bd95-a7b9-419d-918a-f5b4e683f058",
    "isActive": false,
    "balance": "$2,099.20",
    "picture": "http://placehold.it/32x32",
    "age": 28,
    "eyeColor": "blue",
    "name": "Mandy Rodriquez",
    "gender": "female",
    "company": "ZOID",
    "email": "mandyrodriquez@zoid.com",
    "phone": "+1 (944) 547-2976",
    "address": "381 Chase Court, Tilleda, Minnesota, 5631",
    "about": "Esse minim non pariatur quis. Duis id ut aute amet. Laboris minim magna id ex non labore occaecat eiusmod. Culpa in exercitation adipisicing anim elit eiusmod amet Lorem nostrud consectetur consequat anim. Irure veniam laboris et ut est amet mollit officia elit ullamco pariatur. Do laboris ea sint laboris consectetur id consectetur proident nostrud cupidatat qui tempor adipisicing. Dolor ut proident ullamco velit occaecat do officia incididunt ea duis.\r\n",
    "registered": "2015-02-12T05:43:18 -01:00",
    "latitude": -47.958768,
    "longitude": -122.944476,
    "tags": [
      "elit",
      "officia",
      "dolore",
      "ea",
      "adipisicing",
      "pariatur",
      "deserunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Blackburn Mays"
      },
      {
        "id": 1,
        "name": "Estelle Aguilar"
      },
      {
        "id": 2,
        "name": "Franks Watts"
      }
    ],
    "greeting": "Hello, Mandy Rodriquez! You have 4 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7ea3e04c999d79f5aa",
    "index": 1,
    "guid": "849a1238-3381-4ffa-abd0-f5eba618af52",
    "isActive": false,
    "balance": "$2,087.22",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "green",
    "name": "Kenya Calderon",
    "gender": "female",
    "company": "AEORA",
    "email": "kenyacalderon@aeora.com",
    "phone": "+1 (957) 409-3060",
    "address": "244 Holt Court, Fairacres, New York, 3895",
    "about": "Dolore officia do sit magna ex minim reprehenderit eiusmod aliqua ad nulla excepteur. Consectetur occaecat in in esse mollit laboris. Nostrud cupidatat commodo enim anim magna consectetur laboris veniam reprehenderit. Sit ullamco elit est aute ut voluptate ex labore ut. Est deserunt veniam labore occaecat proident. Esse excepteur ad proident deserunt commodo.\r\n",
    "registered": "2015-04-02T04:25:55 -02:00",
    "latitude": -48.994103,
    "longitude": 140.083458,
    "tags": [
      "adipisicing",
      "ad",
      "nostrud",
      "mollit",
      "proident",
      "elit",
      "esse"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Holden Obrien"
      },
      {
        "id": 1,
        "name": "Mays Ratliff"
      },
      {
        "id": 2,
        "name": "Althea Green"
      }
    ],
    "greeting": "Hello, Kenya Calderon! You have 2 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7eb1676978d6717f62",
    "index": 2,
    "guid": "64442276-e35e-4c11-b30e-7407054c847f",
    "isActive": false,
    "balance": "$1,826.67",
    "picture": "http://placehold.it/32x32",
    "age": 40,
    "eyeColor": "blue",
    "name": "Turner Bradshaw",
    "gender": "male",
    "company": "DIGIFAD",
    "email": "turnerbradshaw@digifad.com",
    "phone": "+1 (866) 402-3098",
    "address": "729 Clymer Street, Coultervillle, District Of Columbia, 8706",
    "about": "Fugiat pariatur ullamco magna laboris excepteur ipsum quis quis. Quis pariatur deserunt eu do laborum elit ad do fugiat incididunt laborum proident mollit irure. Velit laboris reprehenderit do nulla ea dolore.\r\n",
    "registered": "2019-02-17T12:00:49 -01:00",
    "latitude": -61.072147,
    "longitude": -103.324679,
    "tags": [
      "sit",
      "consequat",
      "aliqua",
      "nostrud",
      "dolore",
      "ad",
      "id"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Rosemarie Chaney"
      },
      {
        "id": 1,
        "name": "Greta Bishop"
      },
      {
        "id": 2,
        "name": "Cathy Thomas"
      }
    ],
    "greeting": "Hello, Turner Bradshaw! You have 3 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e7d5c59c7aec8bbc6",
    "index": 3,
    "guid": "65cb2d6f-87e8-4b9c-815a-7afa70d5c82d",
    "isActive": false,
    "balance": "$2,548.33",
    "picture": "http://placehold.it/32x32",
    "age": 37,
    "eyeColor": "green",
    "name": "Edwards Miranda",
    "gender": "male",
    "company": "MEDMEX",
    "email": "edwardsmiranda@medmex.com",
    "phone": "+1 (935) 563-3391",
    "address": "954 Furman Avenue, Cedarville, Iowa, 4864",
    "about": "Cillum qui mollit ipsum proident duis est esse aliqua culpa aliquip. Eu nisi cupidatat duis non ea commodo occaecat minim sunt. Duis sunt magna laboris aliqua officia cupidatat esse irure fugiat magna. Veniam esse sit exercitation est eiusmod anim labore. Consequat consectetur commodo culpa culpa nisi tempor nulla. Quis duis proident deserunt adipisicing laborum do eu officia pariatur. Est exercitation aliquip aliquip proident ut duis qui anim pariatur magna fugiat esse ex dolor.\r\n",
    "registered": "2019-01-01T12:53:09 -01:00",
    "latitude": 30.993342,
    "longitude": -129.787686,
    "tags": [
      "quis",
      "ut",
      "aliqua",
      "exercitation",
      "minim",
      "enim",
      "mollit"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Mckay Powers"
      },
      {
        "id": 1,
        "name": "Blair Rocha"
      },
      {
        "id": 2,
        "name": "Benita Talley"
      }
    ],
    "greeting": "Hello, Edwards Miranda! You have 4 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7eba9a438d42ad6b90",
    "index": 4,
    "guid": "666d0d84-a8a5-4b47-89ee-24d1fd221b56",
    "isActive": false,
    "balance": "$2,060.64",
    "picture": "http://placehold.it/32x32",
    "age": 36,
    "eyeColor": "blue",
    "name": "Shana Bartlett",
    "gender": "female",
    "company": "DANCITY",
    "email": "shanabartlett@dancity.com",
    "phone": "+1 (947) 405-3886",
    "address": "338 Orient Avenue, Blodgett, Virginia, 8041",
    "about": "Velit sint dolor irure aliquip aute. Mollit ea aliquip labore tempor mollit in duis et elit occaecat nulla enim ullamco. Aliquip ut eiusmod qui dolor occaecat ex incididunt nulla. Ad amet ex culpa enim veniam velit consectetur. Velit sint est ex amet sint sit do deserunt reprehenderit Lorem eu ut. Occaecat duis pariatur irure reprehenderit id irure fugiat voluptate nisi aliquip labore et.\r\n",
    "registered": "2019-10-19T09:46:48 -02:00",
    "latitude": 15.131421,
    "longitude": 126.273605,
    "tags": [
      "labore",
      "id",
      "fugiat",
      "proident",
      "eiusmod",
      "in",
      "dolore"
    ],
    "friends": [
      {
        "id": 0,
        "name": "West Barker"
      },
      {
        "id": 1,
        "name": "Puckett Collier"
      },
      {
        "id": 2,
        "name": "Terra Carson"
      }
    ],
    "greeting": "Hello, Shana Bartlett! You have 1 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e2b70c9bd2f08f669",
    "index": 5,
    "guid": "b15f3191-d663-4bf6-acdc-07ce0438622d",
    "isActive": false,
    "balance": "$1,638.98",
    "picture": "http://placehold.it/32x32",
    "age": 26,
    "eyeColor": "green",
    "name": "Nadia Williams",
    "gender": "female",
    "company": "ZOXY",
    "email": "nadiawilliams@zoxy.com",
    "phone": "+1 (883) 554-2794",
    "address": "938 Martense Street, Valle, Utah, 5541",
    "about": "Officia eiusmod quis aute deserunt consequat tempor labore officia officia non eu ut. Qui esse aute Lorem minim ipsum laborum sunt officia dolor enim eu do. Veniam dolore ut consequat ex laboris excepteur consequat ea ut minim. Sunt occaecat aliquip id do quis ut dolor amet. Id tempor duis sint dolor dolore ipsum labore ullamco pariatur sunt et cupidatat consectetur.\r\n",
    "registered": "2016-03-25T05:47:08 -01:00",
    "latitude": -34.785955,
    "longitude": -124.629158,
    "tags": [
      "in",
      "consequat",
      "sunt",
      "pariatur",
      "in",
      "eiusmod",
      "in"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Jacobson Kent"
      },
      {
        "id": 1,
        "name": "Florence Stevens"
      },
      {
        "id": 2,
        "name": "Letha Noble"
      }
    ],
    "greeting": "Hello, Nadia Williams! You have 4 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5fd1fa7e5518fe54602ae5d3",
    "index": 6,
    "guid": "db2ddcf2-246b-4e9a-a42b-7305a0f74e6a",
    "isActive": true,
    "balance": "$2,209.84",
    "picture": "http://placehold.it/32x32",
    "age": 28,
    "eyeColor": "brown",
    "name": "Robin Price",
    "gender": "female",
    "company": "DAIDO",
    "email": "robinprice@daido.com",
    "phone": "+1 (931) 418-2091",
    "address": "149 Royce Street, Welch, Alabama, 6557",
    "about": "Duis labore consequat pariatur ipsum Lorem. Occaecat ut minim adipisicing dolore magna proident anim. Exercitation occaecat qui fugiat duis cillum fugiat ipsum ad. Est tempor anim magna elit est laborum minim do officia consectetur ad consequat duis consequat.\r\n",
    "registered": "2014-08-15T03:04:50 -02:00",
    "latitude": -42.412007,
    "longitude": -113.077899,
    "tags": [
      "consectetur",
      "velit",
      "eu",
      "reprehenderit",
      "culpa",
      "velit",
      "consectetur"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Hudson Merrill"
      },
      {
        "id": 1,
        "name": "Heath Serrano"
      },
      {
        "id": 2,
        "name": "Gonzalez Workman"
      }
    ],
    "greeting": "Hello, Robin Price! You have 8 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7e5511d317a85c10ae",
    "index": 7,
    "guid": "9422c07a-2bf8-41e5-a22f-bd4df3ae27bd",
    "isActive": true,
    "balance": "$2,884.30",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "green",
    "name": "Bonnie Herman",
    "gender": "female",
    "company": "IMMUNICS",
    "email": "bonnieherman@immunics.com",
    "phone": "+1 (900) 564-3697",
    "address": "931 Perry Place, Crown, American Samoa, 2861",
    "about": "Occaecat mollit ullamco consectetur non aute fugiat ad exercitation velit aliqua occaecat excepteur mollit. Velit aute labore voluptate deserunt non minim esse proident excepteur eiusmod ex. Adipisicing aute sunt sit voluptate ipsum adipisicing deserunt aliqua laboris voluptate sit aliquip. Esse cillum veniam ullamco reprehenderit esse laboris occaecat. Consequat Lorem enim consectetur id deserunt.\r\n",
    "registered": "2016-10-05T08:58:37 -02:00",
    "latitude": 70.392157,
    "longitude": 51.756576,
    "tags": [
      "duis",
      "deserunt",
      "proident",
      "laborum",
      "eu",
      "anim",
      "qui"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Bridges Cote"
      },
      {
        "id": 1,
        "name": "Osborn Kidd"
      },
      {
        "id": 2,
        "name": "Contreras Wall"
      }
    ],
    "greeting": "Hello, Bonnie Herman! You have 2 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e36da714729f30271",
    "index": 8,
    "guid": "f9727fd2-e2ec-4d96-bdce-fb4de5d429d8",
    "isActive": false,
    "balance": "$1,168.86",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "green",
    "name": "Darcy Weber",
    "gender": "female",
    "company": "COMCUR",
    "email": "darcyweber@comcur.com",
    "phone": "+1 (829) 586-3085",
    "address": "384 Baycliff Terrace, Harold, South Dakota, 8684",
    "about": "Ex nisi ullamco ea esse. Commodo in labore tempor est adipisicing aliqua pariatur velit irure ipsum. Enim velit ea officia officia ad aliquip irure Lorem. Sunt ullamco amet laboris in aliquip.\r\n",
    "registered": "2017-03-20T08:27:59 -01:00",
    "latitude": -81.910668,
    "longitude": -83.086467,
    "tags": [
      "excepteur",
      "duis",
      "irure",
      "exercitation",
      "quis",
      "do",
      "occaecat"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Snyder Waters"
      },
      {
        "id": 1,
        "name": "Benjamin Nash"
      },
      {
        "id": 2,
        "name": "Garner Hooper"
      }
    ],
    "greeting": "Hello, Darcy Weber! You have 5 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e2bc24208323fb01a",
    "index": 9,
    "guid": "1388737f-d427-449d-93c6-4d933e833ab5",
    "isActive": true,
    "balance": "$3,468.37",
    "picture": "http://placehold.it/32x32",
    "age": 23,
    "eyeColor": "blue",
    "name": "Sonya Cohen",
    "gender": "female",
    "company": "YOGASM",
    "email": "sonyacohen@yogasm.com",
    "phone": "+1 (840) 414-2253",
    "address": "572 Lake Place, Hessville, Puerto Rico, 9837",
    "about": "Eiusmod exercitation proident excepteur sunt et et ut. Excepteur sint esse aute Lorem est quis esse eu commodo sunt aliquip. Minim enim ad exercitation ea culpa sint sit dolor proident do dolor occaecat. Eiusmod excepteur consequat irure deserunt dolore consequat proident nisi sint ex dolore amet et consectetur. Elit ex est est occaecat laborum commodo esse pariatur irure nostrud nulla velit magna.\r\n",
    "registered": "2020-02-26T11:29:46 -01:00",
    "latitude": 7.074558,
    "longitude": 65.884986,
    "tags": [
      "nisi",
      "exercitation",
      "non",
      "non",
      "deserunt",
      "fugiat",
      "sit"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Alford House"
      },
      {
        "id": 1,
        "name": "Sharlene Alvarado"
      },
      {
        "id": 2,
        "name": "Terrell Gould"
      }
    ],
    "greeting": "Hello, Sonya Cohen! You have 2 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5fd1fa7e401551bbc53bfb66",
    "index": 10,
    "guid": "336f0357-e6bc-4ff5-84fd-57a2096a83d7",
    "isActive": false,
    "balance": "$1,063.96",
    "picture": "http://placehold.it/32x32",
    "age": 33,
    "eyeColor": "green",
    "name": "Murphy Whitney",
    "gender": "male",
    "company": "COMSTRUCT",
    "email": "murphywhitney@comstruct.com",
    "phone": "+1 (949) 502-3237",
    "address": "119 Sheffield Avenue, Drytown, Maine, 1986",
    "about": "Nulla sunt ullamco enim reprehenderit exercitation consectetur amet exercitation fugiat amet. Do veniam veniam sit ullamco aliquip culpa consequat fugiat sit consequat. Est duis consequat occaecat et.\r\n",
    "registered": "2020-12-04T03:20:29 -01:00",
    "latitude": 59.725137,
    "longitude": 90.17506,
    "tags": [
      "quis",
      "enim",
      "aliquip",
      "irure",
      "qui",
      "culpa",
      "consectetur"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Marian Oliver"
      },
      {
        "id": 1,
        "name": "Iva Jacobson"
      },
      {
        "id": 2,
        "name": "Hebert Carney"
      }
    ],
    "greeting": "Hello, Murphy Whitney! You have 2 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e353e59d77bce7836",
    "index": 11,
    "guid": "afe65b36-27a4-48e3-ad02-5c5c538cf665",
    "isActive": false,
    "balance": "$3,459.84",
    "picture": "http://placehold.it/32x32",
    "age": 23,
    "eyeColor": "green",
    "name": "Sadie Holloway",
    "gender": "female",
    "company": "SOPRANO",
    "email": "sadieholloway@soprano.com",
    "phone": "+1 (804) 574-2807",
    "address": "166 Bowne Street, Eggertsville, Alaska, 2220",
    "about": "Dolore cillum magna nostrud magna fugiat sint nulla nostrud labore. Officia do veniam Lorem nostrud mollit magna ut eiusmod tempor quis excepteur minim. Laboris adipisicing aliquip consequat ut labore incididunt voluptate ad sint nostrud cupidatat aliqua Lorem qui.\r\n",
    "registered": "2018-04-09T05:03:03 -02:00",
    "latitude": -40.84345,
    "longitude": -171.494485,
    "tags": [
      "ex",
      "magna",
      "do",
      "esse",
      "dolore",
      "aute",
      "deserunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Blanche Mckenzie"
      },
      {
        "id": 1,
        "name": "Slater Bennett"
      },
      {
        "id": 2,
        "name": "Bridgette Gillespie"
      }
    ],
    "greeting": "Hello, Sadie Holloway! You have 2 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e9e188c20a2918b17",
    "index": 12,
    "guid": "820f8ca1-ae80-42d3-acb7-581b78c0309f",
    "isActive": false,
    "balance": "$2,626.93",
    "picture": "http://placehold.it/32x32",
    "age": 33,
    "eyeColor": "green",
    "name": "Cathleen Smith",
    "gender": "female",
    "company": "PETICULAR",
    "email": "cathleensmith@peticular.com",
    "phone": "+1 (893) 418-3026",
    "address": "568 Elm Avenue, Fillmore, Tennessee, 3842",
    "about": "Proident consequat aliqua proident proident. Pariatur ex incididunt ex nisi excepteur nostrud sint adipisicing laboris eiusmod. Id sunt nostrud incididunt ad quis officia proident dolor labore nostrud. Sunt sint aliqua aute id ea dolore ipsum dolore culpa pariatur qui Lorem esse voluptate. Qui ut proident aliquip sunt nisi sint amet aliquip ea enim ut adipisicing exercitation. Eiusmod magna consectetur aliquip reprehenderit aliqua do labore amet est. Cillum non aliquip consectetur sint non nisi.\r\n",
    "registered": "2014-06-21T01:21:54 -02:00",
    "latitude": 7.749434,
    "longitude": -23.149009,
    "tags": [
      "anim",
      "minim",
      "ipsum",
      "reprehenderit",
      "labore",
      "laborum",
      "aliqua"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Cherry Kennedy"
      },
      {
        "id": 1,
        "name": "Tammie Pace"
      },
      {
        "id": 2,
        "name": "Jacqueline Key"
      }
    ],
    "greeting": "Hello, Cathleen Smith! You have 7 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e50e98b48588f55dc",
    "index": 13,
    "guid": "8c02905d-2dca-40ac-9d86-2fafa0d4377b",
    "isActive": true,
    "balance": "$1,008.24",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "brown",
    "name": "Hanson Douglas",
    "gender": "male",
    "company": "WARETEL",
    "email": "hansondouglas@waretel.com",
    "phone": "+1 (945) 544-2343",
    "address": "262 Fleet Place, Urbana, Nebraska, 7608",
    "about": "Et est ad minim exercitation. Minim sit culpa excepteur magna dolore voluptate non consectetur veniam. Exercitation deserunt laborum aliquip consequat dolor minim excepteur laboris consectetur sit. Voluptate aliquip cillum pariatur minim proident consequat proident ullamco id magna. Consectetur reprehenderit ex occaecat minim exercitation id reprehenderit nulla. Et in duis labore magna nulla commodo. Labore reprehenderit do veniam veniam dolor ad non amet exercitation nisi sunt cillum dolore dolore.\r\n",
    "registered": "2016-08-21T08:21:42 -02:00",
    "latitude": 9.208136,
    "longitude": 178.538657,
    "tags": [
      "deserunt",
      "aliquip",
      "eu",
      "cillum",
      "proident",
      "eiusmod",
      "ex"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Wilcox Watkins"
      },
      {
        "id": 1,
        "name": "Randi Tran"
      },
      {
        "id": 2,
        "name": "Marci Langley"
      }
    ],
    "greeting": "Hello, Hanson Douglas! You have 1 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5fd1fa7e16e389133b858de2",
    "index": 14,
    "guid": "c6de41ad-0d11-4daa-b4ce-0524580e915b",
    "isActive": false,
    "balance": "$3,540.75",
    "picture": "http://placehold.it/32x32",
    "age": 40,
    "eyeColor": "brown",
    "name": "Carrie Jenkins",
    "gender": "female",
    "company": "ISOLOGIA",
    "email": "carriejenkins@isologia.com",
    "phone": "+1 (996) 427-2504",
    "address": "236 Randolph Street, Austinburg, Nevada, 5283",
    "about": "Ex qui ea fugiat aute anim in officia et. Ut qui cupidatat laborum irure eiusmod esse elit anim do eu labore dolor. Id laborum voluptate sint consequat adipisicing consequat veniam. Esse id labore exercitation cupidatat mollit duis pariatur culpa. Laboris eiusmod labore dolor do culpa incididunt velit in voluptate.\r\n",
    "registered": "2016-12-06T09:13:07 -01:00",
    "latitude": -44.452943,
    "longitude": 27.172302,
    "tags": [
      "fugiat",
      "exercitation",
      "Lorem",
      "magna",
      "Lorem",
      "dolore",
      "Lorem"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Everett Sykes"
      },
      {
        "id": 1,
        "name": "Sheila Knox"
      },
      {
        "id": 2,
        "name": "Heidi Petty"
      }
    ],
    "greeting": "Hello, Carrie Jenkins! You have 10 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7e423758245b3903dc",
    "index": 15,
    "guid": "543863d8-07c0-4787-8dd2-c0d554fd1f0a",
    "isActive": false,
    "balance": "$1,796.88",
    "picture": "http://placehold.it/32x32",
    "age": 35,
    "eyeColor": "green",
    "name": "Webb Barber",
    "gender": "male",
    "company": "POOCHIES",
    "email": "webbbarber@poochies.com",
    "phone": "+1 (875) 486-3660",
    "address": "684 Oakland Place, Cowiche, Federated States Of Micronesia, 6925",
    "about": "Incididunt esse velit amet elit nulla ipsum dolore ut laboris. Magna mollit nulla Lorem incididunt ex velit consequat irure incididunt. Duis exercitation id ad laboris nisi id in irure mollit deserunt labore exercitation. Dolore ut officia excepteur eiusmod ex laborum Lorem dolore.\r\n",
    "registered": "2018-07-07T09:04:16 -02:00",
    "latitude": 20.697968,
    "longitude": 160.387689,
    "tags": [
      "labore",
      "nostrud",
      "ipsum",
      "ea",
      "irure",
      "et",
      "et"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Emma Stone"
      },
      {
        "id": 1,
        "name": "Rosa Phillips"
      },
      {
        "id": 2,
        "name": "Esther Baldwin"
      }
    ],
    "greeting": "Hello, Webb Barber! You have 6 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7e73a7abebbce5a704",
    "index": 16,
    "guid": "dfc92dfa-6b4c-423b-a52d-8eb1caa126e9",
    "isActive": true,
    "balance": "$2,002.06",
    "picture": "http://placehold.it/32x32",
    "age": 27,
    "eyeColor": "green",
    "name": "Lina Dillon",
    "gender": "female",
    "company": "NIKUDA",
    "email": "linadillon@nikuda.com",
    "phone": "+1 (991) 581-2750",
    "address": "181 Anna Court, Calvary, Colorado, 9637",
    "about": "Pariatur aliqua occaecat labore exercitation in in sit. Irure laboris deserunt irure dolore culpa nisi in laborum aute ut quis non dolore. Esse dolore eiusmod cillum velit laboris ea est deserunt occaecat minim consectetur occaecat ut.\r\n",
    "registered": "2014-03-04T01:18:38 -01:00",
    "latitude": -15.719522,
    "longitude": -85.885673,
    "tags": [
      "consectetur",
      "minim",
      "nisi",
      "mollit",
      "est",
      "duis",
      "id"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Brennan Orr"
      },
      {
        "id": 1,
        "name": "Cameron Norris"
      },
      {
        "id": 2,
        "name": "Frazier Higgins"
      }
    ],
    "greeting": "Hello, Lina Dillon! You have 6 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5fd1fa7e8c54cd4c830f0499",
    "index": 17,
    "guid": "46231420-d91b-4e99-bc7a-f53165abde22",
    "isActive": true,
    "balance": "$3,959.74",
    "picture": "http://placehold.it/32x32",
    "age": 23,
    "eyeColor": "green",
    "name": "Landry Russell",
    "gender": "male",
    "company": "ECOLIGHT",
    "email": "landryrussell@ecolight.com",
    "phone": "+1 (932) 539-3304",
    "address": "906 Clarkson Avenue, Lookingglass, Northern Mariana Islands, 8898",
    "about": "Exercitation do eu consectetur anim elit Lorem mollit non cillum dolore velit labore nisi et. Incididunt magna reprehenderit fugiat officia est officia. Minim exercitation exercitation dolore amet Lorem commodo. Deserunt labore est Lorem labore ipsum qui non. Ipsum consequat excepteur adipisicing ea in laborum irure eu magna laborum ipsum magna fugiat tempor. Cupidatat irure cupidatat excepteur voluptate commodo et minim eiusmod.\r\n",
    "registered": "2015-09-18T09:22:18 -02:00",
    "latitude": 64.774821,
    "longitude": 33.563433,
    "tags": [
      "duis",
      "est",
      "proident",
      "aute",
      "nostrud",
      "tempor",
      "ullamco"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Roy Battle"
      },
      {
        "id": 1,
        "name": "Jillian Campos"
      },
      {
        "id": 2,
        "name": "Chen Sloan"
      }
    ],
    "greeting": "Hello, Landry Russell! You have 10 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5fd1fa7e25d5446150a97f55",
    "index": 18,
    "guid": "480d5ccc-1fe5-4e4a-9a5b-243022dddca5",
    "isActive": true,
    "balance": "$3,045.62",
    "picture": "http://placehold.it/32x32",
    "age": 31,
    "eyeColor": "brown",
    "name": "Marissa Shepard",
    "gender": "female",
    "company": "MOREGANIC",
    "email": "marissashepard@moreganic.com",
    "phone": "+1 (974) 435-2810",
    "address": "505 Quentin Road, Carlos, Maryland, 8750",
    "about": "Eu occaecat voluptate ex veniam. Adipisicing tempor amet in laboris do. Aute qui sit pariatur ipsum mollit est dolor commodo ut consectetur. Ex commodo amet eiusmod laboris proident.\r\n",
    "registered": "2016-05-15T06:12:04 -02:00",
    "latitude": -63.157916,
    "longitude": 77.385033,
    "tags": [
      "et",
      "occaecat",
      "sint",
      "ea",
      "occaecat",
      "mollit",
      "culpa"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Talley Benton"
      },
      {
        "id": 1,
        "name": "Day Alvarez"
      },
      {
        "id": 2,
        "name": "Love Nelson"
      }
    ],
    "greeting": "Hello, Marissa Shepard! You have 6 unread messages.",
    "favoriteFruit": "apple"
  }
]`))

type Request struct {
	Protocol      string
	Host          string
	RemoteAddress string
	Method        string
	Path          string
	Headers       http.Header
}

// GetRequest example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Success 200 {string} string	"ok"
// @Router /request/{some_id} [get]
func GetRequest(c echo.Context) error {
	req := c.Request()
	data := Request{
		Protocol:      req.Proto,
		Host:          req.Host,
		RemoteAddress: req.RemoteAddr,
		Method:        req.Method,
		Path:          req.URL.Path,
		Headers:       req.Header,
	}
	return c.JSONPretty(http.StatusOK, data, "    ")
}

func GetSwagger(c echo.Context) error {
	data, err := swag.ReadDoc()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, data)
}

// GetDemoSwag example
// @Summary get random json file
// @Description get json file
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /demoswag [get]
func GetDemoSwag(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, demoBody, "    ")
}
