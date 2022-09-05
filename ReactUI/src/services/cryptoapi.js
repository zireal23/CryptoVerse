import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
const cryptoapiHeaders = {
  "x-rapidapi-host": "coinranking1.p.rapidapi.com",
   "x-rapidapi-key": process.env.REACT_APP_RAPIDAPI_COIN_RANKING_API_KEY,
};

const createRequest = (url) => ({ url, headers: cryptoapiHeaders });

const baseUrl = "https://coinranking1.p.rapidapi.com";
export const cryptoApi = createApi({
  reducerPath: "cryptoApi",
  baseQuery: fetchBaseQuery({ baseUrl }),
  endpoints: (builder) => ({
    getCryptos: builder.query({
      query: (count) => createRequest(`/coins?limit=${count}`),
    }),
    getExchanges: builder.query({
      query: () => createRequest(`/exchanges`),
    }),
    getCryptoDetails: builder.query({
      query: (coinId) => createRequest(`/coin/${coinId}`),
    }),
    getCryptoHistory: builder.query({
      query: ({ coinId, timePeriod }) =>
        createRequest(`/coin/${coinId}/history?timeperiod=${timePeriod}`),
    }),
  }),
});  

export const {
    useGetCryptosQuery,
    useGetCryptoDetailsQuery,
    useGetCryptoHistoryQuery,
    useGetExchangesQuery,
} = cryptoApi;

// var options = {
//   method: "GET",
//   url: "https://coinranking1.p.rapidapi.com/coins",
//   params: {
//     referenceCurrencyUuid: "yhjMzLPhuIDl",
//     timePeriod: "24h",
//     tiers: "1",
//     orderBy: "marketCap",
//     orderDirection: "desc",
//     limit: "50",
//     offset: "0",
//   },
//   headers: {
//     "x-rapidapi-host": "coinranking1.p.rapidapi.com",
//     "x-rapidapi-key": "81b82b588fmshf1cd64975bd20acp10f561jsn0606cfe22d88",
//   },
// };
