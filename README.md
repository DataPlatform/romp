A Romp of Otters
====

## Description

Romp is an environment aware report host and data visualization wrapper. It enables users to configure data sources and inject presentation logic in the form of a mustache.js template. Because the presentation layer is just html/css/js it is possible to create highly custom reports and visualizations.

## Components

##### Restful API

 - Resources include:

     - /sesion
     - /report (Main interface to reports)
     - /report/{report_id}/queries (Set, test, and run queries)
     - /database_connection (Interface for prepopulating database connection options)

 - Use "materialized views" in postgres to cache query data in the database layer. This should be both fast and and preserve database access controls.
 - Minimal state management and no access control. (Access to data is inhereted from the data layer and not messed with.)

##### Report editor

 - Looking at angular.js for interface implementation and ace code editor for textareas that hold source code
 - Inject named SQL select statements that return aggregated data. (This data will be presented as json to the mustache.js template.)
 - Code editor for presentation layer. User is presented with 5 "files" each in a tab:

    - template.html (mustache template)
    - template.css
    - template.js
    - email.css

 - Functionality for previewing report
 - Functionality for scheduling when data should be refreshed
 - Functionality for scheduling email delivery of reports

##### Management routines

 - Report Runner: Something cron like that matches when reports are supposed to be run (against source query logic)
 - Report delivery: Similar service for sending reports via email.
