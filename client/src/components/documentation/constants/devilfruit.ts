import { DocTableDataBody } from '../Documentation';

export const DEVIL_FRUIT_SCHEMA_BODY: DocTableDataBody[] = [
    {
        key: 'avatarSrc',
        type: 'String',
        description: "URL to devil fruit's avatar image.",
    },
    {
        key: 'currentOwner',
        type: 'String',
        description: 'Current owner of the devil fruit',
    },
    {
        key: 'description',
        type: 'String',
        description: 'Brief description of devil fruit.',
    },
    {
        key: 'englishName',
        type: 'String',
        description: "Character's name in English.",
    },
    {
        key: 'id',
        type: 'Int',
        description: 'Unique identifier for devil fruit.',
    },
    {
        key: 'japaneseName',
        type: 'String',
        description: "Devil fruit's name in Japanese.",
    },
    {
        key: 'meaning',
        type: 'String',
        description: "Meaning of the devil fruit's name",
    },
    {
        key: 'previousOwner',
        type: 'String',
        description: 'Previous owner of the devil fruit.',
    },
    {
        key: 'type',
        type: 'String',
        description:
            'Type or category of the devil fruit (eg. Paramecia, Zoan, Logia).',
    },
    {
        key: 'usageDebut',
        type: 'String',
        description: 'First appearance of character.',
    },
];

export const GET_SINGLE_DEVIL_FRUIT_SAMPLE_QUERY = `{
  devilFruit(id: 112) {
    avatarSrc
    currentOwner
    description
    englishName
    id
    meaning
    previousOwner
    type
  }
}`;

export const GET_SINGLE_DEVIL_FRUIT_SAMPLE_RESPONSE = `{
  "data": {
    "devilFruit": {
      "avatarSrc": "https://onepieceql.s3.amazonaws.com/Mochi_Mochi_no_Mi.jpg",
      "currentOwner": "Charlotte Katakuri",
      "description": "The Mochi Mochi no Mi is a Special Paramecia-type Devil Fruit that allows the user to create, control, and transform into mochi. It was eaten by Charlotte Katakuri.",
      "englishName": "Mochi-Mochi Fruit",
      "id": 112,
      "meaning": "Mochi; Springy",
      "previousOwner": "N/A",
      "type": "Special Paramecia"
    }
  }
}`;

export const GET_ALL_DEVIL_FRUITS_SAMPLE_QUERY = `{
  devilFruits(filter: {"page": 1}) {
    info {
      count
      pages
      next
      prev
    }
    results {
      avatarSrc
      currentOwner
      description
      englishName
      id
      meaning
      type
    }
  }
}`;

export const GET_ALL_DEVIL_FRUITS_SAMPLE_RESPONSE = `{
  "data": {
    "characters": {
      "info": {
        "count": 146,
        "next": 2,
        "pages": 4,
        "prev": null
      },
      "results": [
        {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Toshi_Toshi_no_Mi.jpg",
          "currentOwner": "Jewelry Bonney",
          "description": "The Toshi Toshi no Mi is a Paramecia-type Devil Fruit that allows the user to manipulate the ages of people or objects, making them an Age Manipulation Human (年齢自在人間, Nenrei Jizai Ningen?). It was eaten by Jewelry Bonney.",
          "englishName": "Age-Age Fruit",
          "id": 83,
          "meaning": "Age; Year",
          "type": "Paramecia"
        },
        {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Ame_Ame_no_Mi.jpg",
          "currentOwner": "Gasparde",
          "description": "The Ame Ame no Mi is a movie-only Logia Devil Fruit that allows the user to create, control, and transform into candy syrup at will. It was eaten by the main antagonist of the fourth movie, Gasparde.",
          "englishName": "Ame Ame Fruit",
          "id": 41,
          "meaning": "Candy",
          "type": "Logia"
        },
      ]
    }
  }
}`;

export const FILTER_DEVIL_FRUITS_SAMPLE_QUERY = `{
  characters(filter: {"type": "logia", "page": 1}) {
    info {
      count
      pages
      next
      prev
    }
    results {
      avatarSrc
      currentOwner
      description
      englishName
      id
      meaning
      type

    }
  }
}`;

export const FILTER_DEVIL_FRUITS_SAMPLE_RESPONSE = `{
  "data": {
    "characters": {
      "info": {
        "count": 16,
        "next": null,
        "pages": 1,
        "prev": null
      },
      "results": [
 {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Ame_Ame_no_Mi.jpg",
          "currentOwner": "Gasparde",
          "description": "The Ame Ame no Mi is a movie-only Logia Devil Fruit that allows the user to create, control, and transform into candy syrup at will. It was eaten by the main antagonist of the fourth movie, Gasparde.",
          "englishName": "Ame Ame Fruit",
          "id": 41,
          "meaning": "Candy",
          "previousOwner": "N/A",
          "type": "Logia"
        },
        {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Yami_Yami_no_Mi.jpg",
          "currentOwner": "Marshall D. Teach",
          "description": "The Yami Yami no Mi is a Logia-type Devil Fruit that allows the user to create, control, and transform into darkness at will, making the user a Darkness Human (闇人間, Yami Ningen?). It was eaten by Marshall D. Teach, also known as Blackbeard, who stole it from Commander Thatch of the Whitebeard Pirates' 4th division after murdering him.",
          "englishName": "Dark-Dark Fruit",
          "id": 31,
          "meaning": "Darkness",
          "previousOwner": "N/A",
          "type": "Logia"
        },
      ]
    }
  }
}`;
