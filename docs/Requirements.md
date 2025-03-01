System Requirements Specification

Event Scheduler Application (Working Title)

10 February 2025

Draft

_Prepared by:_

Caleb Fringer

Version History

| Change | Version | Date |
| --- | --- | --- |
| Minimum viable product specs | 0.1 | 2/10/25 |
|     |     |     |
|     |     |     |
|     |     |     |
|     |     |     |

**Introduction**

Stakeholders

- Myself (Caleb Fringer)
- DND groups trying to agree on a time to meet for DND
- WoW groups trying to decide on a raid time
- Friend groups trying to decide on a day to go see a movie
- Potential employers and recruiters who want to see my ability to build full stack web applications

**Functional Requirements**

1. The system provides an event organizer with a web interface to create events.
    1. The events page has inputs for the meeting duration, potential dates, and range of hours for which users can select availability.
    2. These inputs constrain the possible dates and times that users can choose their availability.
    3. The interface provides the organizer with inputs for participants phone or emails addresses.
        1. Once the event parameters are selected, the system will send out invitations to participants on this list.
    4. The events page has controls to cancel an event.
2. When an invitation is opened, the system presents invited participants with a web interface to input their availability as a range of possible dates and time ranges.
3. After all participants have input their availability, the system will calculate possible meeting times, notify users, and present participants with a voting web interface.
    1. The voting interface will allow users to vote for the proposed meeting times, and display a waiting screen after casting a vote.
    2. Once all votes are cast, the system will notify users of the winner and provide calendar invite links for Google Calendar, iCalendar, and Outlook.
4. Users will be authenticated with a combination of either email or phone number plus a password, and this data will be stored securely in a database.
    1. Passwords will be hashed before being stored.
5. The system will provide all authenticated users with a home page where they can view events that they are participants in or organizers of.
6. The system will save each event details, including participants, their contact information, proposed meeting times, votes, and availability inputs in the database.
