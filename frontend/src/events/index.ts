import { EventsOn } from '../../wailsjs/runtime'
import useTheme from '../modules/themes/utils/useTheme'
export function bind() {
    EventsOn("change-theme", (name: string) => {
        useTheme(name);
    })
}
