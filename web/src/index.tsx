import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import { ThemeProvider, CSSReset, ColorModeProvider, DarkMode } from '@chakra-ui/core'

import { odinTheme } from './theme'
import App from './App'
import * as serviceWorker from './serviceWorker'
import { RecoilRoot } from 'recoil'

ReactDOM.render(
	<React.StrictMode>
		<BrowserRouter>
			<ThemeProvider theme={odinTheme}>
				<CSSReset />
				<ColorModeProvider value="dark">
					<DarkMode>
						<RecoilRoot>
							<App />
						</RecoilRoot>
					</DarkMode>
				</ColorModeProvider>
			</ThemeProvider>
		</BrowserRouter>
	</React.StrictMode>,
	document.getElementById('root'),
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
