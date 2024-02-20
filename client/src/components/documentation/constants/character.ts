import { DocTableDataBody } from '../Documentation';

export const CHARACTER_SCHEMA_BODY: DocTableDataBody[] = [
    {
        key: 'affiliations',
        type: 'String',
        description: 'Organizations or groups associated with character.',
    },
    {
        key: 'age',
        type: 'String',
        description: 'Age of character.',
    },
    {
        key: 'avatarSrc',
        type: 'String',
        description: "URL to character's avatar image.",
    },
    {
        key: 'birthday',
        type: 'String',
        description: 'Date of birth of character.',
    },
    {
        key: 'bloodType',
        type: 'String',
        description: 'Blood type of character.',
    },
    {
        key: 'bounty',
        type: 'String',
        description:
            'Sum of berries awarded for the capture and/or killing of the character.',
    },

    {
        key: 'debut',
        type: 'String',
        description: 'First appearance of character.',
    },
    {
        key: 'description',
        type: 'String',
        description: 'Brief description of character.',
    },
    {
        key: 'devilFruitName',
        type: 'String',
        description: 'Name of any Devil Fruit character possesses.',
    },
    {
        key: 'englishName',
        type: 'String',
        description: "Character's name in English.",
    },
    {
        key: 'id',
        type: 'Int',
        description: 'Unique identifier for character.',
    },
    {
        key: 'japaneseName',
        type: 'String',
        description: "Character's name in Japanese.",
    },
    {
        key: 'origin',
        type: 'String',
        description: 'Place or region where character originates from.',
    },
];

export const GET_SINGLE_CHARACTER_SAMPLE_QUERY = `{
  character(id: 800) {
    affiliations
    avatarSrc
    birthday
    bloodType
    description
    englishName
  }
}`;

export const GET_SINGLE_CHARACTER_SAMPLE_RESPONSE = `{
  "data": {
    "character": {
      "affiliations": "Straw Hat Pirates",
      "avatarSrc": "https://onepieceql.s3.amazonaws.com/Monkey_D._Luffy_Anime_Post_Timeskip.jpg",
      "birthday": "May 5th (Children's Day)",
      "bloodType": "F",
      "description": "Monkey D. Luffy is the founder and captain of the increasingly infamous and powerful Straw Hat Pirates",
      "englishName": "Monkey D. Luffy"
    }
  }
}`;

export const GET_ALL_CHARACTERS_SAMPLE_QUERY = `{
  characters(filter: {"page": 1}) {
    info {
      count
      pages
      next
      prev
    }
    results {
      affiliations
      avatarSrc
      description
      devilFruitName
      englishName
      id
      origin
    }
  }
}`;

export const GET_ALL_CHARACTERS_SAMPLE_RESPONSE = `{
  "data": {
    "characters": {
      "info": {
        "count": 1345,
        "next": 2,
        "pages": 34,
        "prev": null
      },
      "results": [
        {
          "affiliations": "Ideo Pirates; Straw Hat Grand Fleet",
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Abdullah_Anime.jpg",
          "description": "Abdullah is a criminal and a former bounty hunter who joined as a gladiator at the Corrida Colosseum to compete for the Mera Mera no Mi, along with his partner Jeet. He is an ally of Luffy during the Dressrosa Arc. After the battle with the Donquixote Pirates at Dressrosa, he joined the XXX Gym Martial Arts Alliance, which later converted to a pirate crew called the Ideo Pirates.",
          "devilFruitName": "N/A",
          "englishName": "Abdullah",
          "id": 15,
          "origin": "N/A"
        },
        {
          "affiliations": "Thriller Bark Pirates (Mysterious Four) (former)",
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Absalom_Anime.jpg",
          "description": "Absalom of the Graveyard was the leader of the Zombie Soldiers and Zombie Generals of Thriller Bark prior to its collapse. He was one of the Mysterious Four of the Thriller Bark Pirates and one of the major antagonists of the Thriller Bark Saga.",
          "devilFruitName": "Suke Suke no Mi",
          "englishName": "Absalom",
          "id": 42,
          "origin": "West Blue"
        },
      ]
    }
  }
}`;

export const FILTER_CHARACTERS_SAMPLE_QUERY = `{
  characters(filter: {"affiliation": "straw hat pirates", "page": 1}) {
    info {
      count
      pages
      next
      prev
    }
    results {
      avatarSrc
      description
      devilFruitName
      englishName
      id
      origin
    }
  }
}`;

export const FILTER_CHARACTERS_SAMPLE_RESPONSE = `{
  "data": {
    "characters": {
      "info": {
        "count": 17,
        "next": null,
        "pages": 1,
        "prev": null
      },
      "results": [
        {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Brook_Anime_Post_Timeskip.jpg",
          "description": "\"Soul King\" Brook is the musician of the Straw Hat Pirates, one of their two swordsmen, and one of the Senior Officers of the Straw Hat Grand Fleet. He is the ninth member of the crew and the eighth to join, doing so at the end of the Thriller Bark Arc.",
          "devilFruitName": "Yomi Yomi no Mi",
          "englishName": "Brook",
          "id": 215,
          "origin": "West Blue"
        },
        {
          "avatarSrc": "https://onepieceql.s3.amazonaws.com/Franky_Anime_Post_Timeskip.jpg",
          "description": "\"Iron Man\" Franky is the shipwright of the Straw Hat Pirates and one of the Senior Officers of the Straw Hat Grand Fleet. He is the crew's eighth member and the seventh to join, doing so at the end of the Post-Enies Lobby Arc.",
          "devilFruitName": "N/A",
          "englishName": "Franky; Frankie (originally)",
          "id": 1085,
          "origin": "South Blue"
        },
      ]
    }
  }
}`;
