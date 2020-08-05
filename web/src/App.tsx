import React, { useState } from 'react'
import { Button } from '@chakra-ui/core'

function App() {
	const [test, setTest] = useState('test1')
	return (
		<div>
			<header>
				<p>
					Edit <code>src/App.js</code> and save to reload {test}
					<br />
					<Button variantColor="teal" variant="solid" onClick={() => setTest('1234')}>
						update test
					</Button>
                    
				</p>
				<a href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
					Learn React 2
				</a>
			</header>
		</div>
	)
}

export default App
