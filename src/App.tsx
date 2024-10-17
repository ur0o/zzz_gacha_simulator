import { Routes, Route } from 'react-router-dom'
import { Home } from "./pages/Home.tsx"
import { NotFound } from "./pages/NotFound"
import { css } from '@emotion/react'
import { Gacha } from './pages/Gacha.tsx';

export default function App() {
  const ContainerStyle = css`
    max-width: 1080px;
  `;

  return <>
    <div css={ContainerStyle}>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/gacha/">
          <Route path=":gachaType/:id" element={<Gacha />}/>
        </Route>
        <Route path="" element={<NotFound />} />
      </Routes>
    </div>
  </>
}
