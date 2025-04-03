/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as SmtpImport } from './routes/smtp'
import { Route as RegisterImport } from './routes/register'
import { Route as ImageImport } from './routes/image'
import { Route as DatabaseImport } from './routes/database'
import { Route as IndexImport } from './routes/index'

// Create/Update Routes

const SmtpRoute = SmtpImport.update({
  id: '/smtp',
  path: '/smtp',
  getParentRoute: () => rootRoute,
} as any)

const RegisterRoute = RegisterImport.update({
  id: '/register',
  path: '/register',
  getParentRoute: () => rootRoute,
} as any)

const ImageRoute = ImageImport.update({
  id: '/image',
  path: '/image',
  getParentRoute: () => rootRoute,
} as any)

const DatabaseRoute = DatabaseImport.update({
  id: '/database',
  path: '/database',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/database': {
      id: '/database'
      path: '/database'
      fullPath: '/database'
      preLoaderRoute: typeof DatabaseImport
      parentRoute: typeof rootRoute
    }
    '/image': {
      id: '/image'
      path: '/image'
      fullPath: '/image'
      preLoaderRoute: typeof ImageImport
      parentRoute: typeof rootRoute
    }
    '/register': {
      id: '/register'
      path: '/register'
      fullPath: '/register'
      preLoaderRoute: typeof RegisterImport
      parentRoute: typeof rootRoute
    }
    '/smtp': {
      id: '/smtp'
      path: '/smtp'
      fullPath: '/smtp'
      preLoaderRoute: typeof SmtpImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/database': typeof DatabaseRoute
  '/image': typeof ImageRoute
  '/register': typeof RegisterRoute
  '/smtp': typeof SmtpRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/database': typeof DatabaseRoute
  '/image': typeof ImageRoute
  '/register': typeof RegisterRoute
  '/smtp': typeof SmtpRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/database': typeof DatabaseRoute
  '/image': typeof ImageRoute
  '/register': typeof RegisterRoute
  '/smtp': typeof SmtpRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths: '/' | '/database' | '/image' | '/register' | '/smtp'
  fileRoutesByTo: FileRoutesByTo
  to: '/' | '/database' | '/image' | '/register' | '/smtp'
  id: '__root__' | '/' | '/database' | '/image' | '/register' | '/smtp'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  DatabaseRoute: typeof DatabaseRoute
  ImageRoute: typeof ImageRoute
  RegisterRoute: typeof RegisterRoute
  SmtpRoute: typeof SmtpRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  DatabaseRoute: DatabaseRoute,
  ImageRoute: ImageRoute,
  RegisterRoute: RegisterRoute,
  SmtpRoute: SmtpRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/database",
        "/image",
        "/register",
        "/smtp"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/database": {
      "filePath": "database.tsx"
    },
    "/image": {
      "filePath": "image.tsx"
    },
    "/register": {
      "filePath": "register.tsx"
    },
    "/smtp": {
      "filePath": "smtp.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
