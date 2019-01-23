import BigNumber from "bignumber.js";

export interface Offer {
  offerIndex: BigNumber;
  price: BigNumber;
  duration: BigNumber;
  active: boolean;
  createTime: Date
}
