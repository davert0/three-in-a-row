Cluster1 - классы, представляющие собой игровое поле:
Field, Cell, GameElement, GameElementsGenerator, GameElementsQueue

Cluster2 - классы, представляющие собой поведение игры:
DeleteRule, AddingBonusRule, UseBonusRule, MatchesSet

Cluster3 - классы, определяющие акторов по отношению к игровому полю
Player, Move, Bonus

Cluster4 - utility, статистика
Statistic, GameLog

Cluster5 - infra классы, отвечающие за инициализацию и хранения стейта
Game, GameFactory, SingleTon, GameWatcher
