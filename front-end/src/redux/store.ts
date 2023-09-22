// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import { configureStore } from '@reduxjs/toolkit'
import staticDataSlice from './applicationDataSlice'
import userSlice from './userDataSlice'

export default configureStore({
  reducer: {
    user: userSlice,
    application: staticDataSlice
  }
})