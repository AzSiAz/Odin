import React from 'react'
import { Link } from 'react-router-dom'
import { routes } from '../../constants'

export const LoginScene: React.FC<{}> = ({}) => {
	return (
		<div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
			<div className="max-w-md w-full">
				<div>
					<h2 className="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
						Sign in to your account
					</h2>
				</div>
				<form className="mt-8" action="#" method="POST">
					<input type="hidden" name="remember" value="true" />
					<div className="rounded-md shadow-sm">
						<div>
							<input
								aria-label="Email address"
								name="email"
								type="email"
								required
								className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5"
								placeholder="Email address"
							/>
						</div>
						<div className="-mt-px">
							<input
								aria-label="Password"
								name="password"
								type="password"
								required
								className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5"
								placeholder="Password"
							/>
						</div>
					</div>

					<div className="mt-6 flex justify-between">
						<div className="text-sm leading-5">
							<a
								href="#"
								className="font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:underline transition ease-in-out duration-150"
							>
								Forgot your password?
							</a>
						</div>
						<div className="text-sm leading-5">
							<Link
								to={routes.signUp}
								className="font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:underline transition ease-in-out duration-150"
							>
								Sign up
							</Link>
						</div>
					</div>

					<div className="mt-6">
						<button
							type="submit"
							className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out"
						>
							<span className="absolute left-0 inset-y-0 flex items-center pl-3">
							</span>
							Sign in
						</button>
					</div>
				</form>
			</div>
		</div>
	)
}
