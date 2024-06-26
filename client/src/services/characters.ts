import { RequestPolicy } from '@urql/core';
import { client } from './client';
import { InfoField } from './pagination';

const CHARACTERS_QUERY = `
    query ($page: Int, $filter: characterFilter) {
      characters(page: $page, filter: $filter) {
        results {
          bloodType
          id
          englishName
          japaneseName
          devilFruitName
          description
          bounty
          debut
          birthday
          avatarSrc
          age
          affiliations
          origin
        }
        info {
          count
          next
          pages
          prev
        }
      }
    }
`;

export type Character = {
    affiliations: string;
    age: string;
    avatarSrc: string;
    birthday: string;
    bloodType: string;
    bounty: string;
    debut: string;
    description: string;
    devilFruitName: string;
    englishName: string;
    id: number;
    japaneseName: string;
    origin: string;
};

type CharactersResultField = Character[];

type getCharactersData = {
    characters: {
        info: InfoField;
        results: CharactersResultField;
    };
};

export type CharacterArgs = {
    filter: {
        origin?: string;
        name?: string;
        devilFruitName?: string;
        bloodType?: string;
        birthday?: string;
        affiliations?: string;
    };
    page: number;
};

export const getCharacters = async (
    args: CharacterArgs,
    requestPolicy: RequestPolicy = 'cache-first'
) => {
    return client
        .query<getCharactersData>(CHARACTERS_QUERY, args, { requestPolicy })
        .toPromise()
        .then((res) => res.data?.characters);
};
