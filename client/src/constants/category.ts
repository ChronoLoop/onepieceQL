export const CATEGORIES = ['Character', 'Devil Fruit'] as const;
export type Category = (typeof CATEGORIES)[number];
