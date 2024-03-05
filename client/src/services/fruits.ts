import { RequestPolicy } from '@urql/core';
import { InfoField } from './pagination';
import { client } from './client';

const DEVIL_FRUITS_QUERY = `
    query($page: Int, $filter: devilFruitFilter)  {
      devilFruits(page: $page, filter: $filter) {
        results {
          avatarSrc
          currentOwner
          description
          englishName
          id
          japaneseName
          meaning
          previousOwner
          type
          usageDebut
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

export type DevilFruit = {
    avatarSrc: string;
    currentOwner: string;
    description: string;
    englishName: string;
    id: string;
    japaneseName: string;
    meaning: string;
    previousOwner: string;
    type: string;
    usageDebut: string;
};

type DevilFruitsResultField = DevilFruit[];

type getDevilFruitsData = {
    devilFruits: {
        info: InfoField;
        results: DevilFruitsResultField;
    };
};

export type DevilFruitArgs = {
    filter: {
        name?: string;
        type?: string;
        currentOwner?: string;
        previousOwner?: string;
    };
    page: number;
};

export const getDevilFruits = async (
    args: DevilFruitArgs,
    requestPolicy: RequestPolicy = 'cache-first'
) => {
    return client
        .query<getDevilFruitsData>(DEVIL_FRUITS_QUERY, args, { requestPolicy })
        .toPromise()
        .then((res) => res.data?.devilFruits);
};
