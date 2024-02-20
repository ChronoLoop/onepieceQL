export const INTRODUCTION_SAMPLE_QUERY = `{
  characters(page: "1", filter: {origin: "East Blue"}) {
    info {
      count
      pages
      next
      prev
    },
  }
}`;

export const INTRODUCTION_SAMPLE_RESPONSE = `{
  "data": {
    "characters": {
      "info": {
        "count": 109,
        "next": 2,
        "pages": 3,
        "prev": null
      }
    }
  }
}`;

export const INTRODUCTION_INFO_SCHEMA = [
    {
        key: 'count',
        type: 'Int',
        description: 'Number of records',
    },
    {
        key: 'pages',
        type: 'Int',
        description: 'Total pages',
    },
    {
        key: 'next',
        type: 'Int',
        description: 'Number of next page (null if it does not exist)',
    },
    {
        key: 'prev',
        type: 'Int',
        description: 'Number of previous page (null if it does not exist)',
    },
];

