import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useMemo, useState } from "react";
import { useParams } from "react-router";

namespace Gacha {
  export type Props = {
    gachaType: "character" | "weapon";
    id: number;
  }

  export type Info = {
    id: number;
    name: string;
    puS: string;
    puA_name_0: string;
    puA_name_1: string;
    startDate: string;
    endDate: string;
  };
}

type Result = {
  name: string,
  rank: string,
}

type State = {
  offsetS: number,
  offsetA: number,
  fixedS: boolean,
  fixedA: boolean,
  gacha_id: number
}


export function Gacha() {
  // const gachaInfo = useMemo<Gacha.Info>(() => fetchGachaInfo(), []);
  const [result, setResult] = useState<Result[]>([]);
  const [history, setHistory] = useState<Result[]>([]);
  // const [gachaState, setGachaState] = useState<State>({
  //   offsetS: 0,
  //   offsetA: 0,
  //   fixedS: false,
  //   fixedA: false,
  //   gacha_id: 1
  // });
  const { gachaType, id } = useParams();

  // function fetchGachaInfo() {
  //   const options: AxiosRequestConfig = {
  //     baseURL: BaseUrl,
  //     url: `/gacha/${gachaType}/${id}`,
  //     method: "get",
  //   }
  //   const res = axios(options)
  //   debugger;
  //   return res;
  // }
  function draw() {
    const options: AxiosRequestConfig = {
      baseURL: BaseUrl,
      url: `/gacha/${gachaType}/${id}/draw`,
      method: "get",
      params: {
        offsetS: 0,
        offsetA: 0,
        fixedS: false,
        fixedA: false,
        gacha_id: 1
      }
    }
    axios(options).then((res: AxiosResponse<Result[]>) => {
      const { data } = res;
      setResult(data)
      setHistory(history.concat(result).slice(0, 90))
      debugger;
    })
  }

  return <div>
    <button onClick={() => draw()}>ガチャる</button>
    { result.length != 0 &&
      <table>
        <thead>
          <th>
            <td>レアリティ</td>
            <td>名前</td>
          </th>
        </thead>
        <tbody>
          {result.map((res, _) => {
            return <tr>
              <td>{res.rank}</td>
              <td>{res.name}</td>
            </tr>
          })}
        </tbody>
      </table>
  }
  </div>
}