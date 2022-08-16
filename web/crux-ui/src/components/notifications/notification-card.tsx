import { DyoCard, DyoCardProps } from '@app/elements/dyo-card'
import { DyoHeading } from '@app/elements/dyo-heading'
import { DyoLabel } from '@app/elements/dyo-label'
import DyoTag from '@app/elements/dyo-tag'
import { NotificationItem } from '@app/models'
import clsx from 'clsx'
import useTranslation from 'next-translate/useTranslation'
import Image from 'next/image'

interface NotificationCardProps extends Omit<DyoCardProps, 'children'> {
  notification: NotificationItem
  onClick?: () => void
}

const NotificationCard = (props: NotificationCardProps) => {
  const { t } = useTranslation('notifications')

  const getDefaultImage = <Image src="/notification.svg" width={17} height={21} alt={t('altNotificationPicture')} />

  return (
    <DyoCard className={clsx(props.className ?? 'p-6', 'flex flex-col')}>
      <div className={clsx('flex flex-row flex-grow', props.onClick ? 'cursor-pointer' : null)}>
        {getDefaultImage}

        <DyoHeading className="text-xl text-bright ml-2 my-auto mr-auto" element="h3" onClick={props.onClick}>
          {props.notification.name}
        </DyoHeading>
      </div>

      <div className="flex wrap my-4">
        <p className="text-light break-all">{props.notification.url}</p>
      </div>

      <div className="flex flex-row flex-grow justify-end">
        <DyoLabel className="mr-4 mt-auto py-0.5 leading-4">
          {t('common:createdByName', { name: props.notification.creator })}
        </DyoLabel>
        <DyoTag color="bg-dyo-turquoise" textColor="text-dyo-turquoise">
          {t(`type.${props.notification.type}`)}
        </DyoTag>
      </div>
    </DyoCard>
  )
}

export default NotificationCard