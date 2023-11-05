import { Pair } from './types'

export function SupportedExtensions() {
    return [
        {
          value: 'atest-store-git',
          key: 'atest-store-git'
        },
        {
          value: 'atest-store-s3',
          key: 'atest-store-s3'
        },
        {
          value: 'atest-store-orm',
          key: 'atest-store-orm'
        }
    ] as Pair[]
}